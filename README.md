## Tesla API Client

### This is a full functioning REST API for the Tesla System.

#### First steps are to get your auth token, see below:
```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("POST", 
                                "https://owner-api.teslamotors.com/oauth/token?grant_type=password&client_id=YOUR_CLIENT_IDI&client_secret=YOUR_CLIENT_SECRET&email=TESLA_EMAIL&password=TESLA_PASSWORD", 
                                nil)
	if err != nil {
		// log err if there are any issues
	}
	resp, err := client.Do(req)
	if err != nil {
		// again, error handling is option for you
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body)) // log the result, which will be your auth token
}
``` 

#### After this, you can create your program and do things like:
```go
package main

import (
	"github.com/Noy/Tesla"
	"log"
	"net/http"
)
// Authentication
func authTesla() *tesla.AuthTesla {
	var id, authToken string
    teslaCar := tesla.AuthTesla{ID: id, AccessToken: authToken}
    return &teslaCar
}

//Then...
// Use this function in your desired ways.. It'll honk your car!
func honk() {
    authTesla().HonkHorn()
}
```

##### Ref: https://www.teslaapi.io/vehicles/state-and-settings