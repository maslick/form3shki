package main

import (
	"fmt"
	"github.com/maslick/form3shki"
	"log"
)

func main() {
	client := form3shki.New()
	list, err := client.List()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, account := range list {
		fmt.Println(account.Attributes.Country)
	}
}
