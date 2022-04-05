package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Ikan" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ajung", spamFilter)
	sayHelloWithFilter("Ikan", spamFilter)
}
