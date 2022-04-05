package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var aju Customer
	aju.Name = "Ajung"
	aju.Address = "Indonesia"
	aju.Age = 30

	aju.sayHi("Setia")
	aju.sayHuuu()

	// fmt.Println(aju.Name)
	// fmt.Println(aju.Address)
	// fmt.Println(aju.Age)

	// set := Customer{
	// 	Name:    "Setia",
	// 	Address: "Cirebon",
	// 	Age:     35,
	// }
	// fmt.Println(set)

	// tia := Customer{"Tia", "Jakarta", 35}
	// fmt.Println(tia)
}
