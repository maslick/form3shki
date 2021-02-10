# form3shki
Golang client library for Form3 API (see [specs](ASSIGNMENT.md))


# Features
* Unit tests
* CI/CD Github Actions
* Example

# Development
```
docker-compose up --build
go test -v -coverprofile=coverage.txt
go tool cover -func=coverage.txt
go tool cover -html=coverage.txt
```