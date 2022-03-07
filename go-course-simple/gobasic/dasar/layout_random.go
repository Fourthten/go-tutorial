package main

import (
	"fmt"
	"math/rand"
	"time"
)

type _student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = _student{
	name:        "setia",
	height:      180.5,
	age:         25,
	isGraduated: false,
	hobbies:     []string{"swim", "play"},
}

func main() {
	fmt.Printf("%b\n", data.age)
	fmt.Printf("%c\n", 1400)
	fmt.Printf("%c\n", 1235)
	fmt.Printf("%d\n", data.age)
	fmt.Printf("%e\n", data.height)
	fmt.Printf("%E\n", data.height)
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.9f\n", data.height)
	fmt.Printf("%.2f\n", data.height)
	fmt.Printf("%.f\n", data.height)
	fmt.Printf("%e\n", 0.123123123)
	fmt.Printf("%f\n", 0.123123123)
	fmt.Printf("%g\n", 0.123123123)
	fmt.Printf("%g\n", 0.12)
	fmt.Printf("%.5g\n", 0.12)
	fmt.Printf("%o\n", data.age)
	fmt.Printf("%p\n", &data.name)
	fmt.Printf("%q\n", `" name \ height "`)
	fmt.Printf("%s\n", data.name)
	fmt.Printf("%t\n", data.isGraduated)
	fmt.Printf("%T\n", data.name)
	fmt.Printf("%T\n", data.height)
	fmt.Printf("%T\n", data.age)
	fmt.Printf("%T\n", data.isGraduated)
	fmt.Printf("%T\n", data.hobbies)
	fmt.Printf("%v\n", data)
	fmt.Printf("%+v\n", data)
	fmt.Printf("%#v\n", data)

	var datas = struct {
		name   string
		height float64
	}{
		name:   "setia",
		height: 180.5,
	}
	fmt.Printf("%#v\n", datas)

	fmt.Printf("%x\n", data.age)
	var d = data.name
	fmt.Printf("%x%x%x\n", d[0], d[1], d[2])
	fmt.Printf("%x\n", d)
	fmt.Printf("%X\n", d)
	fmt.Printf("%%\n")

	rand.Seed(10)
	fmt.Println("random ke-1:", rand.Int())
	fmt.Println("random ke-2:", rand.Int())

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Int())
	fmt.Println(rand.Int())
	fmt.Println(rand.Float32())
	fmt.Println(rand.Uint32())
	fmt.Println(rand.Intn(2))

	randomString(10)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
