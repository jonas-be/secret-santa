package email

import (
	"net/smtp"
)

type Auth struct {
	Sender         string
	Password       string
	SmtpServer     string
	SmtpServerPort string
}

func (auth Auth) getAuth() smtp.Auth {
	return smtp.PlainAuth("", auth.Sender, auth.Password, auth.getServerAddress())
}

func (auth Auth) getServerAddress() string {
	return auth.SmtpServer + ":" + auth.SmtpServerPort
}
