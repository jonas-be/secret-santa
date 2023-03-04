package config

type EmailConfig struct {
	Subject string `yaml:"subject"`
	Content string `yaml:"content"`
}
