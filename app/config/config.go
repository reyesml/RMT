package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database struct {
		DbId string `yaml:"dbId"`
	} `yaml:"database"`
	Server struct {
		Port uint `yaml:"port"`
	} `yaml:"server"`
	Session struct {
		SigningSecret string `yaml:"signing_secret"`
	} `yaml:"session"`
}

func LoadConfig(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config: %w", err)
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to decode config: %w", err)
	}
	return cfg, nil
}
