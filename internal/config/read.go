package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func Read() (Config, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	fileContents, err := os.ReadFile(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading the contents of config file: %w", err)
	}

	cfg := Config{}
	err = json.Unmarshal(fileContents, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Error decoding JSON: %w", err)
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error when reading HOME variable: %w", err)
	}
	file := home + "/" + configFileName

	return file, nil
}
