package user

import "fmt"

type User struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

func (user User) sendMail(content string) {
	fmt.Printf("Sending Email to %v\nWith content:\n%v", user.Email, content)
}
