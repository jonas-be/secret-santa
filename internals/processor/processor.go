package processor

import (
	"fmt"
	"math/rand"
	"secret-santa/internals/config"
	"secret-santa/internals/user"
)

type Processor struct {
	Config *config.Config
}

var userMapping UserMap = make(map[user.User]user.User)

func (processor Processor) Process() {
	i := 0
	for {
		i++
		processor.generateRandomMapping()
		if processor.checkValidity() {
			break
		}
	}
	fmt.Printf(" + + + Result + + + \n\n")
	fmt.Println(userMapping)
	fmt.Printf("Succed after %v tries", i)
}

func (processor Processor) generateRandomMapping() {
	userPool := make([]user.User, len(processor.Config.Users))
	copy(userPool, processor.Config.Users)

	for _, u := range processor.Config.Users {
		poolIndex := rand.Intn(len(userPool))
		userMapping[u] = userPool[poolIndex]
		userPool = append(userPool[:poolIndex], userPool[poolIndex+1:]...)
	}
}

func (processor Processor) checkValidity() bool {
	for k, v := range userMapping {
		if k == v {
			return false
		}
	}
	return processor.checkRuleSet()
}

func (processor Processor) checkRuleSet() bool {
	for _, ban := range processor.Config.ComboBans {
		for _, u1 := range ban.Combo {
			for _, u2 := range ban.Combo {
				if !processor.checkComboRule(u1, u2) {
					return false
				}
			}
		}
	}
	return true
}

func (processor Processor) checkComboRule(u1 string, u2 string) bool {
	for k, v := range userMapping {
		if k.Name == u1 && v.Name == u2 {
			return false
		}
	}
	return true
}
