package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Ajung"
	sayHelloTo(firstName, "Setia")
	sayHelloTo("Tia", "Agung")
}
