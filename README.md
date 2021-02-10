# =form3shki=
Golang client library for Form3 API (see [specs](ASSIGNMENT.md))


## :rocket: Features
* CI/CD (Github Actions)
* Unit tests
* Example demo
* Production ready

## :lollipop: Usage
```shell
$ mdkir test && cd test
$ go mod init example.com/test
$ go get github.com/maslick/form3shki@v0.4.0
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
	config.SetBaseUrl("http://localhost:8080")

	// initialize client
	client,_ := form3shki.NewClientWithConfig(config)

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
```

## :bulb: Development
```shell
$ docker-compose up --build
$ go test -v -coverprofile=coverage.txt
$ go tool cover -func=coverage.txt
$ go tool cover -html=coverage.txt
```
