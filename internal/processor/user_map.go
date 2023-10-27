package processor

import (
	"fmt"
	"secret-santa/internal/config"
	"secret-santa/internal/email"
)

type UserMap map[config.User]config.User

func (userMap UserMap) String() string {
	userMappingListString := ""
	for k, v := range userMap {
		userMappingListString += fmt.Sprintf("<p><b>%v => %v</b>   <i>(%v hat %v gezogen)</i></p>\n", k.Name, v.Name, k.Name, v.Name)
	}
	return userMappingListString
}

func (userMap UserMap) sendMails(conf config.EmailConfig, summaryMail string) {
	auth, _ := email.GetAuthFromEnv()

	userMap.sendSummaryMail(conf, summaryMail, auth)

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

func (userMap UserMap) sendSummaryMail(conf config.EmailConfig, summaryMail string, auth email.Auth) {
	if summaryMail == "" {
		return
	}
	mail := email.Email{
		Auth:     auth,
		Receiver: summaryMail,
		Content: email.Content{
			Subject: conf.Subject + " | Summary",
			Body:    userMap.String(),
		},
	}
	err := mail.SendMail()
	if err != nil {
		fmt.Println(err)
		return
	}
}
