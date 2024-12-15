package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AccountJSON string
	Parent      string
}

func LoadConfig() *Config {
	var config Config
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Cannot decode config JSON: %v", err)
	}

	return &config
}
