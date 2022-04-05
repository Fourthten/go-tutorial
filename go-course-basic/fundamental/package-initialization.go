package main

import (
	"fmt"
	"go-course-basic/fundamental/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
