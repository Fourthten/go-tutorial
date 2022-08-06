package main

import (
	"fmt"
	"go-course-basic/fundamental/helper"
)

func main() {
	helper.SayHello("Ajung")
	// helper.sayGoodbye("Ajung") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
