package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("vim-go")
	Anything(2.44)
	Anything("Setia")
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap["name"] = "Setia"
	mymap["age"] = 4
	fmt.Println(mymap)
}
