# =form3shki=
Golang client library for Form3 API (see [specs](ASSIGNMENT.md))

[![Build](https://github.com/maslick/form3shki/workflows/Build/badge.svg)](https://github.com/maslick/form3shki/actions)

## :rocket: Features
* CRUD operations on Accounts resource (Create, Fetch, List, Delete)
* List operation supports pagination  
* Integration tests
* CI pipeline (Github Actions)
* Golint
* Example demo app

## :lollipop: Usage ([see example](example/main.go))
```shell
$ mdkir test && cd test
$ go mod init example.com/test
$ go env -w GOPRIVATE=github.com/maslick
$ go get github.com/maslick/form3shki@v0.5.0
$ touch main.go
```

```go
// main.go
package main

import (
	"fmt"
	"github.com/maslick/form3shki"
	"log"
)

func main() {
	// create configuration
	config := form3shki.NewConfig()
	config.SetBaseURL("http://localhost:8080")

	// initialize client
	client,_ := form3shki.NewClientWithConfig(config)

	// fetch/list accounts (first page, take 10 items)
	list, err := client.List(0, 10)

	if err != nil {
		log.Fatal(err.Error())
	}

	// print
	for _, account := range list {
		fmt.Println(account.ID)
	}
}
```

## :bulb: Development
```shell
$ docker-compose up --build
$ go test -v -coverprofile=coverage.txt
$ go tool cover -html=coverage.txt
$ go tool cover -func=coverage.txt
github.com/maslick/form3shki/client.go:27:	init			90.0%
github.com/maslick/form3shki/client.go:49:	NewClient		100.0%
github.com/maslick/form3shki/client.go:61:	NewClientWithConfig	100.0%
github.com/maslick/form3shki/client.go:68:	Create			76.9%
github.com/maslick/form3shki/client.go:92:	Fetch			84.6%
github.com/maslick/form3shki/client.go:119:	List			81.8%
github.com/maslick/form3shki/client.go:138:	Delete			75.0%
github.com/maslick/form3shki/client.go:160:	getEnv			0.0%
github.com/maslick/form3shki/config.go:9:	NewConfig		100.0%
github.com/maslick/form3shki/config.go:14:	BaseURL			100.0%
github.com/maslick/form3shki/config.go:19:	SetBaseURL		0.0%
total:						(statements)		85.7%
```
