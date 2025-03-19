package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Port string `json:"port"`
	} `json:"server"`
	Database struct {
		Host     string `json:"host"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Port     string `json:"port"`
	} `json:"database"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.Open("dev.json")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Failed to decode config JSON: %v", err)
	}
}
