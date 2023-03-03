package processor

import (
	"fmt"
	"secret-santa/internals/user"
)

type UserMap map[user.User]user.User

func (userMap UserMap) String() string {
	userMappingListString := ""
	for k, v := range userMap {
		userMappingListString += fmt.Sprintf("%v :: %v\n", k.Name, v.Name)
	}
	return userMappingListString
}
