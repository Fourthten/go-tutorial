package main

import (
	"fmt"
	"strings"
)

// type person struct { name string; age int; }
type person struct {
	address string
	age     int
}

type student struct {
	name  string
	grade int
}

type students struct {
	name  string
	grade int
	age   int
	person
}

type detailStudent struct {
	persons struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

// type Animal = struct { name string; age int }
type Animal struct {
	name string `tag_name`
	age  int    `tag_age`
}

// type Number = int
// var num Number = 4
type Pet = Animal

func main() {
	// Struct, Pointer
	var numberA int = 4
	var numberB *int = &numberA
	fmt.Println(numberA, &numberA, *numberB, numberB)
	numberA = 5
	fmt.Println(numberA, &numberA, *numberB, numberB)

	change(&numberA, 10)
	fmt.Println(numberA)

	var s1 student
	s1.name = "Setia"
	s1.grade = 4
	fmt.Println(s1, s1.name)
	var s2 = student{}
	s2.name = "Ajung"
	s2.grade = 4
	var s3 = student{"Setia", 4}
	var s4 = student{name: "Ajung"}
	var s5 = student{grade: 4, name: "Ajung"}
	var s6 *student = &s1
	fmt.Println(s1.name, s2.name, s3.name, s4.name, s5.name, s6.name)
	s6.name = "Ajung"
	fmt.Println(s1.name, s6.name)

	var s7 = students{}
	s7.age = 20
	s7.person.age = 21
	s7.address = "here"
	var p1 = person{address: "here", age: 20}
	var s8 = students{person: p1, grade: 20}
	fmt.Println(s7.address, s7.age, s7.person.age)
	fmt.Println(s8)

	// var p1 = struct { name string; age int } { age: 20, name: "Setia" }
	// var p1 = struct { name string; age int } { "Setia", 20 }
	var s9 = struct {
		person
		grade int
	}{}
	s9.person = person{"here", 20}
	s9.grade = 4
	var s10 = struct {
		person
		grade int
	}{
		person: person{"here", 21},
		grade:  4,
	}
	fmt.Println(s9, s10)

	var allStudents = []person{
		{address: "Asia", age: 20},
		{address: "Java", age: 20},
	}
	for _, student := range allStudents {
		fmt.Println(student.address, student.age)
	}

	var allStudents2 = []struct {
		person
		grade int
	}{
		{person: person{"here", 21}, grade: 4},
		{person: person{"java", 21}, grade: 4},
	}
	for _, student := range allStudents2 {
		fmt.Println(student)
	}

	var student1 struct {
		person
		grade int
	}
	student1.person = person{"here", 21}
	student1.grade = 4
	var student2 struct {
		grade int
	}
	var student3 = struct {
		grade int
	}{
		4,
	}
	fmt.Println(student1, student2, student3)

	var animal = Animal{"Cat", 4}
	animal2 := Animal{"Cat", 4}
	var pet = Pet{"Fish", 4}
	pet2 := Pet{"Fish", 4}
	fmt.Println(animal, pet)
	fmt.Println(Pet(animal2))
	fmt.Println(Animal(pet2))

	// Method
	var s11 = student{"Ajung", 20}
	var name = s11.getNameAt(2)
	s11.sayHello()
	fmt.Println("name: ", name)

	fmt.Println("student: ", s11.name)
	s11.changeName1("Setia")
	fmt.Println("student: ", s11.name)
	s11.changeName2("Fourth")
	fmt.Println("student: ", s11.name)
	var s12 = &student{"Ajung", 21}
	s12.sayHello()
	fmt.Println(strings.Split("Ajung Setia", " "))
}

func change(original *int, value int) {
	*original = value
}

func (s student) sayHello() {
	fmt.Println("Hello", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s student) changeName1(name string) {
	fmt.Println("name: ", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("name: ", name)
	s.name = name
}
