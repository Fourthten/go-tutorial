package main

import "fmt"

func main() {
	name := "Ajung"
	counter := 0

	increment := func() {
		name := "Setia"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
