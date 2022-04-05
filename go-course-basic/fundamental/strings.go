package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Ajung Setia", "Ajung"))
	fmt.Println(strings.Contains("Ajung Setia", "Agung"))

	fmt.Println(strings.Split("Ajung Setia Agung", " "))

	fmt.Println(strings.ToLower("Ajung Setia Agung"))
	fmt.Println(strings.ToUpper("Ajung Setia Agung"))
	fmt.Println(strings.ToTitle("ajung setia agung"))

	fmt.Println(strings.Trim("      Ajung Setia     ", " "))
	fmt.Println(strings.ReplaceAll("Ajung Agung Ajung", "Ajung", "Setia"))
}
