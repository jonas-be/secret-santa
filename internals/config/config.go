package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"secret-santa/internals/user"
)

type Config struct {
	Users     []user.User `yaml:"users"`
	ComboBans []ComboBan  `yaml:"comboBans"`
}

func (conf Config) String() string {
	userListString := ""
	comboBanListString := ""
	for _, u := range conf.Users {
		userListString += fmt.Sprintf("\tName: %v\n\t      %v\n", u.Name, u.Email)
	}
	for i, u := range conf.ComboBans {
		usernameList := ""
		for _, username := range u.Combo {
			usernameList += fmt.Sprintf("\t\t%v\n", username)
		}
		comboBanListString += fmt.Sprintf("\t%v:\n%v", i+1, usernameList)
	}
	return fmt.Sprintf("Users[%v]:\n%v\nComboBans[%v]:\n%v", len(conf.Users), userListString, len(conf.ComboBans), comboBanListString)
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
