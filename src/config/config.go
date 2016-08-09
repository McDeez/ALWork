package config

import (
	"encoding/json"
	"log"
	"os"
)

// Simple config struct for all parameters
// TargetHost: The host we will be proxying to: http://url.goes.here.com[:port]
type Config struct {
	TargetHost    string
	ListeningPort string
}

var MyConfig *Config = nil

func LoadConfig() *Config {
	if MyConfig == nil {
		file, _ := os.Open("config.json")
		decoder := json.NewDecoder(file)
		MyConfig = &Config{}
		err := decoder.Decode(&MyConfig)
		if err != nil {
			// Default settings if no config.json was found
			log.Println("Error loading config: using defaults")
			MyConfig.ListeningPort = ":8080"
			MyConfig.TargetHost = "http://localhost:3000"
		}
	}
	return MyConfig
}
