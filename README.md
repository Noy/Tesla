## Tesla API Client

### This is a full functioning REST API for the Tesla System.

#### First steps are to get your auth token, see below:
```go
package main

import (
    "context"
    tesla "github.com/Noy/Tesla"
    "log"
)

func main() {
	authToken, err := tesla.RetrieveToken(context.Background(), "tesla@email.com", "teslaPassword")
	if err != nil {
		// error handling is option for you
	}
	log.Println(authToken) // log the result, which will be your auth token
}
``` 

#### Then, you need to retrieve your car's ID:
```go
package main

import (
	"github.com/Noy/Tesla"
)
// Authentication
func getCarId() int64 {
    authToken := "abc123"
    // Use the auth token from before to authenticate with AuthTesla
    auth := tesla.AuthTesla{AccessToken: authToken}
    id := auth.ListVehicles()[0].ID
    return id
}
```

#### Finally, you can retrieve data about your car (or even perform actions!):
```go
package main

import (
	"github.com/Noy/Tesla"
)
// Authentication
func authTesla() *tesla.AuthTesla {
	var id, authToken string
    // Use the auth token and ID from before to authenticate with AuthTesla
    teslaCar := tesla.AuthTesla{ID: id, AccessToken: authToken}
    return &teslaCar
}

//Then...
// Use this function in your desired ways.. It'll honk your car!
func honk() {
    authTesla().HonkHorn()
}
```

##### Ref: https://www.teslaapi.io/
##### Special thanks to: https://github.com/jsgoecke/tesla for new auth