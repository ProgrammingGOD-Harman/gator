package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

func Read() (Config, error) {
	fullpath, err := getConfigfilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullpath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	cfg := Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigfilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(homeDir, configFileName)

	return fullPath, nil
}

func write(cfg Config) error {
	fullpath, err := getConfigfilePath()

	if err != nil {
		return err
	}

	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)

	if err != nil {
		return err
	}

	return nil
}
