package tesla

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/manifoldco/promptui"
	"golang.org/x/oauth2"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
	This is purely for authentication! This should be the first thing you run before integrating this API. See README.
*/
// Hats off to https://github.com/jsgoecke/tesla for this.

func RetrieveToken(ctx context.Context, email, password string) (string, error) {
	verifier, challenge, err := pkce()
	if err != nil {
		return "", fmt.Errorf("pkce: %w", err)
	}

	c := &oauth2.Config{
		ClientID:     "ownerapi",
		ClientSecret: "",
		RedirectURL:  "https://auth.tesla.com/void/callback",
		Scopes:       []string{"openid email offline_access"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://auth.tesla.com/oauth2/v3/authorize",
			TokenURL: "https://auth.tesla.com/oauth2/v3/token",
		},
	}

	hc := &http.Client{
		Transport: &Transport{RoundTripper: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}},
	}

	code, err := (&Auth{
		AuthURL: c.AuthCodeURL(state(), oauth2.AccessTypeOffline,
			oauth2.SetAuthURLParam("code_challenge", challenge),
			oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		),
		SelectDevice: selectDevice,
		Client:       hc,
	}).Do(ctx, email, password)
	if err != nil {
		return "", err
	}

	ctx = context.WithValue(ctx, oauth2.HTTPClient, hc)

	t, err := c.Exchange(ctx, code,
		oauth2.SetAuthURLParam("code_verifier", verifier),
	)
	if err != nil {
		return "", fmt.Errorf("exchange: %w", err)
	}

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "\t")
	return t.AccessToken, nil
}

func userAgent() string {
	const prefixBytes = 6
	var buf [prefixBytes + 8]byte
	if _, err := io.ReadFull(rand.Reader, buf[:]); err != nil {
		// I know.
		panic(err)
	}

	var b strings.Builder
	e := base32.NewEncoder(base32.StdEncoding.WithPadding(base32.NoPadding), &b)
	e.Write(buf[:prefixBytes])
	e.Close()
	b.WriteRune('/')
	b.Write(strconv.AppendUint(nil, binary.BigEndian.Uint64(buf[prefixBytes:]), 10))
	return b.String()
}

type Transport struct {
	http.RoundTripper
	userAgent     string
	userAgentTime time.Time
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if now := time.Now(); t.userAgent == "" || now.Sub(t.userAgentTime) > 6*time.Hour {
		t.userAgent = userAgent()
		t.userAgentTime = now
	}
	for _, h := range []struct{ k, v string }{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"User-Agent", t.userAgent},
	} {
		if _, ok := req.Header[h.k]; ok {
			continue
		}
		req.Header.Set(h.k, h.v)
	}
	res, err := t.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if strings.EqualFold(res.Header.Get("Content-Encoding"), "gzip") {
		res.Body = &gzipReader{body: res.Body}
	}
	return res, err
}

type gzipReader struct {
	body io.ReadCloser
	zr   *gzip.Reader
	zerr error
}

func (gz *gzipReader) Read(p []byte) (n int, err error) {
	if gz.zr == nil {
		if gz.zerr == nil {
			gz.zr, gz.zerr = gzip.NewReader(gz.body)
		}
		if gz.zerr != nil {
			return 0, gz.zerr
		}
	}
	return gz.zr.Read(p)
}

func (gz *gzipReader) Close() error {
	return gz.body.Close()
}

func state() string {
	var b [9]byte
	if _, err := io.ReadFull(rand.Reader, b[:]); err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(b[:])
}

// https://www.oauth.com/oauth2-servers/pkce/
func pkce() (verifier, challenge string, err error) {
	var p [87]byte
	if _, err := io.ReadFull(rand.Reader, p[:]); err != nil {
		return "", "", fmt.Errorf("rand read full: %w", err)
	}
	verifier = base64.RawURLEncoding.EncodeToString(p[:])
	b := sha256.Sum256([]byte(challenge))
	challenge = base64.RawURLEncoding.EncodeToString(b[:])
	return verifier, challenge, nil
}

func selectDevice(ctx context.Context, devices []Device) (d Device, passcode string, err error) {
	var i int
	if len(devices) > 1 {
		var err error
		i, _, err = (&promptui.Select{
			Label:   "Device",
			Items:   devices,
			Pointer: promptui.PipeCursor,
		}).Run()
		if err != nil {
			return Device{}, "", fmt.Errorf("select device: %w", err)
		}
	}
	d = devices[i]

	passcode, err = (&promptui.Prompt{
		Label:   "Passcode",
		Pointer: promptui.PipeCursor,
		Validate: func(s string) error {
			if len(s) != 6 {
				return errors.New("len(s) != 6")
			}
			return nil
		},
	}).Run()
	if err != nil {
		return Device{}, "", err
	}
	return d, passcode, nil
}

type Device struct {
	DispatchRequired bool      `json:"dispatchRequired"`
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	FactorType       string    `json:"factorType"`
	FactorProvider   string    `json:"factorProvider"`
	SecurityLevel    int       `json:"securityLevel"`
	Activated        time.Time `json:"activatedAt"`
	Updated          time.Time `json:"updatedAt"`
}

type Auth struct {
	Client       *http.Client
	AuthURL      string
	SelectDevice func(ctx context.Context, devices []Device) (d Device, passcode string, err error)
}

func (a *Auth) Do(ctx context.Context, username, password string) (code string, err error) {
	if a.Client == nil {
		a.Client = &http.Client{
			Transport: &http.Transport{
				Dial: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		}
	}

	if a.Client.Jar == nil {
		var err error
		a.Client.Jar, err = cookiejar.New(nil)
		if err != nil {
			return "", fmt.Errorf("new cookie jar: %w", err)
		}
	}

	cr := a.Client.CheckRedirect
	a.Client.CheckRedirect = func(*http.Request, []*http.Request) error {
		return http.ErrUseLastResponse
	}
	defer func() { a.Client.CheckRedirect = cr }()

	res, v, err := a.login(ctx, username, password)
	if err != nil {
		return "", fmt.Errorf("login: %w", err)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
	case http.StatusFound:
		return codeFromResponse(res)
	default:
		return "", fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	transactionID := v.Get("transaction_id")

	devices, err := a.listDevices(ctx, transactionID)
	if err != nil {
		return "", fmt.Errorf("list devices: %w", err)
	}

	if len(devices) == 0 {
		return "", errors.New("no devices")
	}

	d, passcode, err := a.SelectDevice(ctx, devices)
	if err != nil {
		return "", fmt.Errorf("select device: %w", err)
	}

	if err := a.verify(ctx, transactionID, d, passcode); err != nil {
		return "", fmt.Errorf("verify: %w", err)
	}
	return a.commit(ctx, transactionID)
}

func (a *Auth) login(ctx context.Context, username, password string) (*http.Response, url.Values, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.AuthURL, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("new request: %w", err)
	}

	res, err := a.Client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("do: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	v := url.Values{
		"identity":   {username},
		"credential": {password},
	}

	d, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("new document: %w", err)
	}

	d.Find("input[type=hidden]").Each(func(_ int, s *goquery.Selection) {
		name, ok := s.Attr("name")
		if !ok {
			return
		}
		value, ok := s.Attr("value")
		if !ok {
			return
		}
		v.Set(name, value)
	})

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, a.AuthURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, nil, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err = a.Client.Do(req)
	return res, v, err
}

func (a *Auth) listDevices(ctx context.Context, transactionID string) ([]Device, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, (&url.URL{
		Scheme:   "https",
		Host:     "auth.tesla.com",
		Path:     "oauth2/v3/authorize/mfa/factors",
		RawQuery: url.Values{"transaction_id": {transactionID}}.Encode(),
	}).String(), nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	res, err := a.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	var out struct {
		Data []Device `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("json decode: %w", err)
	}
	return out.Data, nil
}

func (a *Auth) verify(ctx context.Context, transactionID string, d Device, passcode string) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(map[string]string{
		"transaction_id": transactionID,
		"factor_id":      d.ID,
		"passcode":       passcode,
	}); err != nil {
		return fmt.Errorf("json encode: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://auth.tesla.com/oauth2/v3/authorize/mfa/verify", &buf)
	if err != nil {
		return fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := a.Client.Do(req)
	if err != nil {
		return fmt.Errorf("do: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var out struct {
		Data struct {
			Approved bool `json:"approved"`
		} `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	if !out.Data.Approved {
		return errors.New("not approved")
	}
	return nil
}

func (a *Auth) commit(ctx context.Context, transactionID string) (code string, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.AuthURL, strings.NewReader(url.Values{
		"transaction_id": {transactionID},
	}.Encode()))
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := a.Client.Do(req)
	if err != nil {
		return "", fmt.Errorf("do: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusFound {
		return "", fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
	return codeFromResponse(res)
}

func codeFromResponse(res *http.Response) (code string, err error) {
	u, err := res.Location()
	if err != nil {
		return "", fmt.Errorf("response location: %w", err)
	}
	return u.Query().Get("code"), nil
}
