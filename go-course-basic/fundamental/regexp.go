package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])s")

	fmt.Println(regex.MatchString("aas"))
	fmt.Println(regex.MatchString("abs"))
	fmt.Println(regex.MatchString("aNs"))

	search := regex.FindAllString("aas aja ads abs ays ani", -1)
	fmt.Println(search)
}
