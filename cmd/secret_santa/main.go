package main

import (
	"fmt"
	"secret-santa/internals/config"
	"secret-santa/internals/processor"
)

func main() {
	conf, err := config.LoadConfig("test.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(conf)

	processor := processor.Processor{Config: conf}
	processor.Process()
}
