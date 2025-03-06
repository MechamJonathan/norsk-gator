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
	return Write(*cfg)
}

func Read() (Config, error) {
	fullpath, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullpath)
	if err != nil {
		return Config{}, err
	}

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func GetConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullpath := filepath.Join(home, configFileName)
	return fullpath, nil
}

func Write(cfg Config) error {
	fullpath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullpath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
