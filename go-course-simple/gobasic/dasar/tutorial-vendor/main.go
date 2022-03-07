package main

import (
	"fmt"

	gubrak "github.com/novalagung/gubrak/v2"
)

func main() {
	fmt.Println(gubrak.RandomInt(10, 20))
}

// go mod init tutorial-vendor
// go get -u repo
// go mod vendor
// >= 1.14
// go run main.go
// go build -o executable
// < 1.14
// go run -mod=vendor main.go
// go build -mod=vendor -o executable

// mac or linux
// ./executable
// windows
// executable.exe
