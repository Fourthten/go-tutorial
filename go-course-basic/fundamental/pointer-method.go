package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	ajung := Man{"Ajung"}
	ajung.Married()

	fmt.Println(ajung.Name)
}
