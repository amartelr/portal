package main

import (
	"encoding/json"

	"log"
	"os"

	"github.com/amartelr/portal/webportal"
)

type configuration struct {
	ServerAddress string `json:"webserver"`
}

func main() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}

	config := new(configuration)
	json.NewDecoder(file).Decode(config)

	log.Println("Connecting ... on addr ", config.ServerAddress)
	webportal.Run(config.ServerAddress)
}
