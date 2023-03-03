package main

import (
	"fmt"
	"secret-santa/internals/config"
	"secret-santa/internals/processor"
)

func main() {
	conf, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println(err)
	}

	processor := processor.Processor{Config: conf}
	processor.Process()
}
