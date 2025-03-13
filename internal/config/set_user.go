package config

import (
	"encoding/json"
	"os"
)

func (c Config) SetUser(user string) error {
	c.CurrentUsername = user
	err := write(c)

	return err
}

func write(cfg Config) error {
	configFile, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	os.WriteFile(configFile, data, 0644)
	return nil
}
