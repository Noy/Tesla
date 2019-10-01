package main

import (
	"bytes"
	"encoding/json"
	"github.com/BurntSushi/toml"
	"log"
)

func main() {
	// this will eventually be a library package, main used for testing
}

type Config struct {
	AccessToken, ID, RefreshToken string
}

func loadConfig() (c Config, err error) {
	_, err = toml.DecodeFile("config.toml", &c)
	if err != nil {
		log.Printf("Could not load config: %v", err.Error())
	}
	return
}

func getConfig() Config {
	cfg, err := loadConfig()
	if err != nil {
		log.Println("Could not open config")
	}
	return cfg
}

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
