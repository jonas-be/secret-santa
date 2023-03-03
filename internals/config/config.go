package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"secret-santa/internals/user"
)

type Config struct {
	Users     []user.User `yaml:"users"`
	ComboBans []ComboBan  `yaml:"forbiddenCombinations"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
