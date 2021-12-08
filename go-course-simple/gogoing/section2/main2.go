package main

import "fmt"

/* arrays */
/*
func todo() {
	// var arr []int
	arr := []int{1, 2, 3, 4}
	// fmt.Println(arr)

	arr2 := []string{"hi", "my", "name"}

	arr2 = append(arr2, "is", "setia", "!")
	fmt.Println(arr, arr2)

}

func main() {
	todo()
}
*/

/* structure encapsulation */
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("driving...")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	c := Car{
		Name:    "agung",
		Age:     1,
		ModelNo: 2,
	}
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}
