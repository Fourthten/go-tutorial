package main

import "fmt"

func main() {
	var name = "Setia"

	if name == "Ajung" {
		fmt.Println("Hello Ajung")
	} else if name == "Agung" {
		fmt.Println("Hello Agung")
	} else if name == "Tia" {
		fmt.Println("Hello Tia")
	} else {
		fmt.Println("Hi, kenalan donk")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
