package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	Account string `yaml:"account"`
	Vault   string `yaml:"vault"`
	Item    string `yaml:"item"`
}

type Config struct {
	Profiles map[string]Profile `yaml:"profiles"`
}

func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(homeDir, ".opssh.yaml"), nil
}

func LoadConfig() (*Config, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{Profiles: make(map[string]Profile)}, nil
		}
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	if config.Profiles == nil {
		config.Profiles = make(map[string]Profile)
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func GetProfile(config *Config, name string) (*Profile, error) {
	profile, exists := config.Profiles[name]
	if !exists {
		return nil, fmt.Errorf("profile '%s' not found", name)
	}
	return &profile, nil
}

func AddProfile(config *Config, name string, profile Profile) error {
	if _, exists := config.Profiles[name]; exists {
		return fmt.Errorf("profile '%s' already exists", name)
	}
	config.Profiles[name] = profile
	return nil
}

func RemoveProfile(config *Config, name string) error {
	if _, exists := config.Profiles[name]; !exists {
		return fmt.Errorf("profile '%s' not found", name)
	}
	delete(config.Profiles, name)
	return nil
}

