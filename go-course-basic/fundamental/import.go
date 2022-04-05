package main

import (
	"fmt"
	"go-course-basic/fundamental/helper"
)

func main() {
	helper.SayHello("Ajung")
	// helper.sayGoodbye("Eko") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
