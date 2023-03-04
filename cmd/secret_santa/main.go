package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"secret-santa/internals/config"
	"secret-santa/internals/email"
	"secret-santa/internals/processor"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conf, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}
	if _, err := email.GetAuthFromEnv(); err != nil {
		fmt.Println(err)
		return
	}

	processor := processor.Processor{Config: conf}
	processor.Process()
}
