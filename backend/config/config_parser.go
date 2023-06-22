// Package config for config details
package config

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/validator.v2"
)

// Configuration struct to hold app configurations
type Configuration struct {
	Port       string     `json:"port"`
	Database   DB         `json:"database"`
	StripeKeys StripeKeys `json:"stripe"`
	Version    string     `json:"version" validate:"nonzero"`
}

// DB struct to hold database file
type DB struct {
	File string `json:"file" validate:"nonzero"`
}

// StripeKeys struct to hold stripe keys
type StripeKeys struct {
	Publisher string `json:"publisher" validate:"nonzero"`
	Secret    string `json:"secret" validate:"nonzero"`
}

// ReadConfFile read configurations of json file
func ReadConfFile(path string) (Configuration, error) {
	config := Configuration{}
	file, err := os.Open(path)
	if err != nil {
		return Configuration{}, fmt.Errorf("failed to open config file: %w", err)
	}

	dec := json.NewDecoder(file)
	if err := dec.Decode(&config); err != nil {
		return Configuration{}, fmt.Errorf("failed to load config: %w", err)
	}

	return config, validator.Validate(config)
}
