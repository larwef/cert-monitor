package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Config holds configuration options for the app
type Config struct {
	Endpoint     string   `json:"endpoint"`
	ClientID     string   `json:"clientId"`
	ClientSecret string   `json:"clientSecret"`
	TokenURL     string   `json:"tokenUrl"`
	Scopes       []string `json:"scopes"`
}

// New reads config from a file and returns a config object
func New(filepath string) *Config {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Exiting application. Error reading config file: %v\n", err)
		os.Exit(1)
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		fmt.Printf("Exiting application. Error parsing config file: %v\n", err)
		os.Exit(1)
	}

	return &config
}
