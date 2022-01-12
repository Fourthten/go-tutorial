package main

import (
	"fmt"
)

func main() {
	// Print text
	fmt.Println("Hello World")
	fmt.Println("Hello", "World!")

	// Variable
	var firstName = "Setia"
	var lastName string
	lastName = "Ajung"
	middleName := "Fourth"
	fmt.Printf("Hello %s %s!\n", firstName, lastName)
	fmt.Println("Hello", firstName, middleName+"!")

	var first, second, third string
	first, second, third = "one", "two", "three"
	var fourth, fifth, sixth = "four", "five", "six"
	seventh, isEighth, ninth, tenth := "seven", true, 9, 10.0
	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, isEighth, ninth, tenth)

	_ = "not used"
	name, _ := "Setia", "not used"
	midName := new(string)
	*midName = "Fourth"
	fmt.Println(name, midName)
	fmt.Println(*midName)

	// Data type
	var valString = ""
	var valInteger = 0
	var valDecimal = 0.0
	var isBool = false
	fmt.Println(valString, valInteger, valDecimal, isBool)

	// Constanta
	const valConst string = "constanta"
	const valConsts = " const"
	fmt.Print(valConst, valConsts, "\n")

	// Operator
	var valArithmetic = (((5 + 6) % 2) * 4) / 2
	var isEqual = (valArithmetic == 10)
	var isNotEqual = (valArithmetic != 10)
	var isSmall = (valArithmetic < 10)
	var isSmallEqual = (valArithmetic <= 10)
	var isBig = (valArithmetic > 10)
	var isBigEqual = (valArithmetic >= 10)
	fmt.Printf("nilai %d (%t) (%t) (%t) (%t) (%t) (%t)\n",
		valArithmetic, isEqual, isNotEqual, isSmall, isSmallEqual, isBig, isBigEqual)

	var left = false
	var right = true
	var isAnd = left && right
	var isOr = left || right
	var isReverse = !left
	fmt.Printf("Logic \t(%t) (%t) (%t)\n", isAnd, isOr, isReverse)

	// Condition
	if valArithmetic == 2 {
		fmt.Println("Two")
	} else if valArithmetic == 4 {
		fmt.Println("Four")
	} else {
		fmt.Println("Nothing")
	}
	if percent := valDecimal / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect\n", percent, "%")
	} else if percent >= 60 {
		fmt.Printf("%.1f%s perfect\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s perfect\n", percent, "%")
	}
	switch valInteger {
	case 100:
		fmt.Printf("Full\n")
	case 50, 60:
		fmt.Printf("Half\n")
	default:
		fmt.Printf("Zero\n")
	}
	switch {
	case valDecimal == 100:
		fmt.Printf("Full\n")
	case (valDecimal < 100), (valDecimal > 50):
		fmt.Printf("Half\n")
	default:
		{
			fmt.Printf("Zero\n")
		}
	}
	switch {
	case valArithmetic == 10:
		fmt.Printf("Full\n")
	case (valArithmetic < 10), (valArithmetic > 1):
		fmt.Printf("Half\n")
		fallthrough // next condition is true
	case valArithmetic < 2:
		fmt.Printf("Low\n")
	default:
		{
			fmt.Printf("Zero\n")
		}
	}
	if valArithmetic == 2 {
		switch valArithmetic {
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("Nothing")
		}
	} else {
		if valArithmetic == 4 {
			fmt.Println("Four")
		} else {
			fmt.Println("Nothing")
		}
	}

	// Looping
	for i := 0; i < 1; i++ {
		fmt.Println("Number", i)
	}
	for valArithmetic <= 4 {
		fmt.Println("Number", valArithmetic)
		valArithmetic++
	}
	for i := 1; i < 4; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 3 {
			break
		}
		fmt.Println("Number", i)
	}
	for i := 1; i <= 2; i++ {
		for j := i; j <= 2; j++ {
			fmt.Print(i, ".", j, " ")
		}
		fmt.Println()
	}
outerLoop:
	for i := 1; i <= 2; i++ {
		for j := i; j <= 2; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(i, ".", j, " ")
		}
		fmt.Println()
	}

	// Array
	var names [2]string
	names[0] = "Setia"
	names[1] = "Ajung"
	var fruits = [2]string{"Mango", "Melon"}
	var numbers = [...]int{1, 2}
	var num = [2][3]int{[3]int{1, 2, 3}, [3]int{1, 2, 3}}
	var nums = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(names[0], names[1])
	fmt.Println("Fruits ", len(fruits), fruits)
	fmt.Println("Number ", len(numbers), numbers)
	fmt.Println("Number ", num)
	fmt.Println("Number ", nums)

	var nameList [2]string
	nameList = [2]string{
		"Setia",
		"Ajung",
	}
	for i := 0; i < len(nameList); i++ {
		fmt.Printf("element %d : %s\n", i, nameList[i])
	}
	for i, name := range nameList {
		fmt.Printf("element %d : %s\n", i, name)
	}
	for _, name := range nameList {
		fmt.Printf("element : %s\n", name)
	}
	for i, _ := range nameList {
		fmt.Print(i)
	}
	for i := range nameList {
		fmt.Print(i)
	}
	var makeFruits = make([]string, 2)
	makeFruits[0] = "Apple"
	makeFruits[1] = "Banana"
	fmt.Println("\n", makeFruits)

	// Slice
	var slaceFruits = []string{"Apple", "Banana", "Orange"}
	var newFruits = slaceFruits[1:2]
	var slaceFruit = append(slaceFruits, "Mango")
	fmt.Println(slaceFruits, len(slaceFruits), cap(slaceFruits))
	fmt.Println(newFruits, len(newFruits), cap(newFruits))
	fmt.Println(slaceFruit, len(slaceFruit), cap(slaceFruit))

	dst := make([]string, 3)
	src := []string{"Apple", "Orange", "Banana"}
	n := copy(dst, src)
	fmt.Println(dst, src, n)
	dst2 := []string{"Potato", "Potato"}
	src2 := []string{"Apple"}
	n2 := copy(dst2, src2)
	fmt.Println(dst2, src2, n2)

	var aSlace = slaceFruits[0:2]
	var bSlace = slaceFruits[0:2:2]
	fmt.Println(aSlace, len(aSlace), cap(aSlace))
	fmt.Println(bSlace, len(bSlace), cap(bSlace))

	// Map
	var chicken1 map[string]int
	chicken1 = map[string]int{}
	chicken1["feb"] = 10
	var chicken2 = map[string]int{"jan": 20}
	var chicken3 = make(map[string]int)
	chicken3["feb"] = 10
	var chicken4 = *new(map[string]int)
	fmt.Println(chicken1, chicken2, chicken3, chicken4)
	for key, val := range chicken1 {
		fmt.Println(key, " :", val)
	}

	var chicken5 = map[string]int{
		"january":  20,
		"february": 10,
	}
	delete(chicken5, "january")
	fmt.Println(len(chicken5), chicken5)

	var value, isExist = chicken1["jan"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item not exist")
	}
	var chickens = []map[string]string{
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken blue", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println(chicken["name"], chicken["gender"])
	}
	var data = []map[string]string{
		{"description": "chicken red", "gender": "male"},
		{"age": "2", "id": "1"},
	}
	fmt.Println(data)
}

/*
	Basic Code
*/
