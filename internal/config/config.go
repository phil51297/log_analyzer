package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type LogConfig struct {
	ID   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func LoadConfig(configPath string) ([]LogConfig, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("impossible de lire le fichier de configuration: %w", err)
	}

	var configs []LogConfig
	if err := json.Unmarshal(data, &configs); err != nil {
		return nil, fmt.Errorf("erreur de parsing JSON: %w", err)
	}

	return configs, nil
}
