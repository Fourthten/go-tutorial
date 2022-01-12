package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	var name = []string{"Setia", "Ajung"}
	printMessage("Hello", name)
	divideNumber(10, 2)

	var diameter float64 = 10
	var area, circumference = calculate(diameter)
	fmt.Printf("luas: %.2f \n", area)
	fmt.Printf("keliling: %.2f \n", circumference)

	var avg = calculate2(1, 2, 3, 4, 5)
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
	var numbers = []int{1, 2, 3, 4, 5}
	var avg2 = calculate2(numbers...)
	var msg2 = fmt.Sprintf("Rata-rata : %.2f", avg2)
	fmt.Println(msg2)

	hobbies("Setia", "Sleep", "Swim")
	var hobby = []string{"Sleep", "Eat"}
	hobbies("Setia", hobby...)

	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}
	var numbers2 = []int{1, 2, 3}
	var min, max = getMinMax(numbers2)
	fmt.Printf("data: %v min: %v max: %v", numbers, min, max)

	// var closure (func (string, int, []string) int)
	// closure = func (a string, b int, c []string) int { }
	var numbers3 = []int{1, 2, 3}
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)
	fmt.Println("original number: ", numbers3)
	fmt.Println("filtered number: ", newNumbers)

	var max1 = 3
	var numbers4 = []int{1, 2, 3}
	var howMany, getNumber = findMax(numbers4, max1)
	var theNumber = getNumber()
	fmt.Println("max: ", max1, "number: ", numbers4, "found: ", howMany, "value: ", theNumber)

	// type FilterCallback func(string) bool
	// func filter(data []string, callback FilterCallback) []string { }
	var data = []string{"Setia", "Ajung"}
	var dataConstain = filter(data, func(each string) bool {
		return strings.Contains(each, "g")
	})
	var dataLength = filter(data, func(each string) bool {
		return len(each) == 5
	})
	fmt.Println("data: ", data)
	fmt.Println("constain g: ", dataConstain)
	fmt.Println("length: ", dataLength)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)

	rand.Seed(time.Now().Unix())
	var randomValue int
	randomValue = randomWithRange(2, 10)
	fmt.Println("Random number: ", randomValue)
}

func randomWithRange(min, max int) int {
	// func nameFunc(A type, B type, C type) returnType
	// func nameFunc(A, B, C type) returnType
	var value = rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

func calculate(d float64) (float64, float64) {
	// func calculate(d float64) (area float64, circumference float64) { return }
	var area = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d
	return area, circumference
}

func calculate2(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func hobbies(name string, hobbies ...string) {
	var hobString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, %s\n", name)
	fmt.Printf("Hobby: %s\n", hobString)
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
	// var function = func() []int { return res }
	// return len(res), function
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
