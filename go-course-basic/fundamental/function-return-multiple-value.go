package main

import "fmt"

func getFullName() (string, string, string) {
	return "Ajung", "Setia", "Agung"
}

func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}
