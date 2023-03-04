package email

import (
	"net/smtp"
	"os"
)

type Auth struct {
	Sender         string
	Password       string
	SmtpServer     string
	SmtpServerPort string
}

type AuthInvalidError struct {
	s string
}

func (e AuthInvalidError) Error() string {
	return e.s
}

func (auth Auth) getAuth() smtp.Auth {
	return smtp.PlainAuth("", auth.Sender, auth.Password, auth.SmtpServer)
}

func (auth Auth) getServerAddress() string {
	return auth.SmtpServer + ":" + auth.SmtpServerPort
}

func (auth Auth) isValid() (bool, error) {
	if auth.Sender == "" {
		return false, AuthInvalidError{"Sender is empty"}
	}
	if auth.Password == "" {
		return false, AuthInvalidError{"Password is empty"}
	}
	if auth.SmtpServer == "" {
		return false, AuthInvalidError{"SmtpServer is empty"}
	}
	if auth.SmtpServerPort == "" {
		return false, AuthInvalidError{"SmtpServerPort is empty"}
	}
	return true, nil
}

func GetAuthFromEnv() (Auth, error) {
	toReturn := Auth{
		os.Getenv("MAIL_SENDER"),
		os.Getenv("MAIL_PW"),
		os.Getenv("MAIL_SMTP_SERVER"),
		os.Getenv("MAIL_SMTP_SERVER_PORT"),
	}
	if valid, err := toReturn.isValid(); !valid {
		return toReturn, err
	}
	return toReturn, nil
}
