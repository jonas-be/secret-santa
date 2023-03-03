package processor

import (
	"fmt"
	"secret-santa/internals/config"
)

type UserMap map[config.User]config.User

func (userMap UserMap) String() string {
	userMappingListString := ""
	for k, v := range userMap {
		userMappingListString += fmt.Sprintf("%v :: %v\n", k.Name, v.Name)
	}
	return userMappingListString
}
