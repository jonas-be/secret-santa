package email

import (
	"net/smtp"
)

type Email struct {
	Auth     Auth
	Receiver string
	Content  Content
}

func (email Email) sendMail() error {
	return smtp.SendMail(
		email.Auth.getServerAddress(),
		email.Auth.getAuth(),
		email.Auth.Sender,
		[]string{email.Receiver},
		email.Content.toByteArray(),
	)
}
