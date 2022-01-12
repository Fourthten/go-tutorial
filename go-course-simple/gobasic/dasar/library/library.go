package library

import "fmt"

type Student struct {
	Name  string
	Grade int
}

var Students = struct {
	Name  string
	Grade int
}{}

func SaysHello() {
	fmt.Println("Hello")
}

func Introduce(name string) {
	fmt.Println("name", name)
}

func SayHi(name string) {
	fmt.Println("Hy")
	Introduce(name)
}

func init() {
	Students.Name = "Ajung"
	Students.Grade = 4
}
