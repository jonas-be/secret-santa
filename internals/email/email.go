package email

import (
	"fmt"
	"net/smtp"
)

type Email struct {
	Auth     Auth
	Receiver string
	Content  Content
}

func (email Email) SendMail() error {
	fmt.Printf("Sennding mail to %v ", email.Receiver)
	err := smtp.SendMail(
		email.Auth.getServerAddress(),
		email.Auth.getAuth(),
		email.Auth.Sender,
		[]string{email.Receiver},
		email.Content.toByteArray(),
	)
	if err != nil {
		fmt.Println(" >>> FAILED")
		return err
	}
	fmt.Println(" >>> DONE")
	return nil
}
