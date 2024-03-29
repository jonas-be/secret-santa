package main

import (
	"flag"
	"fmt"
	"log"
	"secret-santa/internal/config"
	"secret-santa/internal/email"
	"secret-santa/internal/processor"

	"github.com/joho/godotenv"
)

var yFlag = flag.Bool("y", false, "Send mails directly")

func main() {
	flag.Parse()

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

	processor := processor.Processor{Config: conf, AskBeforeSending: !*yFlag}
	processor.Process()
}
