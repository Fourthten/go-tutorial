package main

import (
	f "fmt"
	"go-course-simple/gobasic/dasar/library"
	. "go-course-simple/gobasic/dasar/library"
)

func main() {
	library.SaysHello()
	library.Introduce("Ajung")
	library.SayHi("Setia")

	var s1 = library.Student{"Ajung", 19}
	f.Println("name ", s1.Name)
	f.Println("grade ", s1.Grade)

	var s2 = Student{"Ajung", 20}
	f.Println("name ", s2.Name)
	f.Println("grade ", s2.Grade)

	f.Printf("Name: %s ", library.Students.Name)
	f.Printf("Grade: %d\n", library.Students.Grade)

	// go run main.go partial.go
	// sayHy("Setia")
}
