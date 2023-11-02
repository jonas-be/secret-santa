package processor

import (
	"fmt"
	"secret-santa/internal/config"
	"secret-santa/internal/email"
)

type UserMap map[config.User]config.User

func (userMap UserMap) getLongetsName() int {
	longestName := 0
	for k := range userMap {
		if len(k.Name) > longestName {
			longestName = len(k.Name)
		}
	}
	return longestName
}

func AddSpaces(s string, totalLen int) string {
	for i := len(s); i < totalLen; i++ {
		s += " "
	}
	return s
}

func (userMap UserMap) String() string {
	maxLen := userMap.getLongetsName()*2 + 5
	userMappingListString := ""
	for k, v := range userMap {
		userMappingListString += fmt.Sprintf(" %v  (%v has drawn %v)\n", AddSpaces(fmt.Sprintf(" %v => %v", k.Name, v.Name), maxLen), k.Name, v.Name)
	}
	return userMappingListString
}

func (userMap UserMap) StringHTML() string {
	userMappingListString := ""
	for k, v := range userMap {
		userMappingListString += fmt.Sprintf("<p><b>%v => %v</b>   <i>(%v has drawn %v)</i></p>\n", k.Name, v.Name, k.Name, v.Name)
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
			Body:    userMap.StringHTML(),
		},
	}
	err := mail.SendMail()
	if err != nil {
		fmt.Println(err)
		return
	}
}
