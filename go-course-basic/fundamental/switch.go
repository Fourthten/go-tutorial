package main

import "fmt"

func main() {

	name := "Setia"

	switch name {
	case "Ajung":
		fmt.Println("Hello Ajung")
		fmt.Println("Hello Ajung")
	case "Tia":
		fmt.Println("Hello Tia")
		fmt.Println("Hello Tia")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
