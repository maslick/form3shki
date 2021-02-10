package main

import (
	"fmt"
	"github.com/maslick/form3shki"
	"log"
)

func main() {
	// create configuration
	config := form3shki.NewConfig()
	config.SetBaseUrl("http://localhost:8080")

	// initialize client
	client, _ := form3shki.NewClientWithConfig(config)

	// fetch/list all accounts
	list, err := client.List()

	if err != nil {
		log.Fatal(err.Error())
	}

	// print
	for _, account := range list {
		fmt.Println(account.Id)
	}
}
