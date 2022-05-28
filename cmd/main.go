package main

import (
	"log"

	client "github.com/Agilen/Client"
)

func main() {
	err := client.Start(nil)
	if err != nil {
		log.Fatal(err)
	}
}
