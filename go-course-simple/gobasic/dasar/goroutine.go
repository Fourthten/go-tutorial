package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go print(5, "Hello")
	print(5, "How are u")
	var input string
	fmt.Scanln(&input)

	// func Scanln(a ...interface{}) (n int, err error)
	var s1, s2, s3 string
	fmt.Scanln(&s1, &s2, &s3)
	fmt.Println(s1, s2, s3)
}
