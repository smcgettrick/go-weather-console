package main

import (
	"encoding/json"
	"os"
)

// Configuration contains values from conf.json
type Configuration struct {
	AppID string `json:"AppId"`
}

// LoadConfiguration loads config data from conf.json
func LoadConfiguration() Configuration {
	var configuration Configuration
	configFile, err := os.Open("conf.json")
	if err != nil {
		panic(err.Error())
	}

	parser := json.NewDecoder(configFile)
	parser.Decode(&configuration)

	configFile.Close()

	return configuration
}
