package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	defer fmt.Println("Hello")
	fmt.Println("Welcome")

	orderSomeFood("burger")
	orderSomeFood("pizza")

	number := 3
	if number == 3 {
		fmt.Println("Hello")
		// defer fmt.Println("Hey")
		func() {
			defer fmt.Println("Hey")
		}()
	}
	fmt.Println("Hy")

	defer fmt.Println("Hello")
	// os.Exit(1)
	fmt.Println("Welcome back")

	// error
	var input string
	fmt.Print("Type number: ")
	fmt.Scanln(&input)
	var number2 int
	var err error
	number2, err = strconv.Atoi(input)
	if err == nil {
		fmt.Println(number2, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}

	// custom error
	var name string
	fmt.Print("Type name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("Hello", name)
	} else {
		// fmt.Println(err.Error())
		panic(err.Error())
		fmt.Println("end")
	}

	// recover
	defer catch()
	var name3 string
	fmt.Print("Type name: ")
	fmt.Scanln(&name3)
	if valid, err := validate(name3); valid {
		fmt.Println("Hello", name3)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred", r)
		} else {
			fmt.Println("Apps running perfectly")
		}
	}()
	// panic("some error happened")

	data := []string{"ajung", "setia", "agung"}
	for _, each := range data {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occurred", each, "| message:", r)
				} else {
					fmt.Println("App running perfectly")
				}
			}()
			panic("some error")
		}()
	}
}

func orderSomeFood(menu string) {
	defer fmt.Println("Thank you,")
	if menu == "pizza" {
		fmt.Print("Nice choice,", " ")
		fmt.Print("Pizza is delicious", " ")
		return
	}

	fmt.Println("Your order", menu)
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot empty")
	}
	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occurred", r)
	} else {
		fmt.Println("App running perfectly")
	}
}
