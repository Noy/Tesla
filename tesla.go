package tesla

import (
	"bytes"
	"encoding/json"
)

type AuthTesla struct {
	ID, AccessToken string
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
