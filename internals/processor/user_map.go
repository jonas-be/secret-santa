package processor

import (
	"fmt"
	"secret-santa/internals/config"
	"secret-santa/internals/email"
)

type UserMap map[config.User]config.User

func (userMap UserMap) String() string {
	userMappingListString := ""
	for k, v := range userMap {
		userMappingListString += fmt.Sprintf("%v :: %v\n", k.Name, v.Name)
	}
	return userMappingListString
}

func (userMap UserMap) sendMails(conf config.EmailConfig) {
	auth, _ := email.GetAuthFromEnv()
	for k, v := range userMap {
		mail := email.Email{
			Auth:     auth,
			Receiver: k.Email,
			Content: email.Content{
				Subject: conf.Subject,
				Body:    fmt.Sprintf(conf.Content, k.Name, v.Name),
			},
		}
		err := mail.SendMail()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
