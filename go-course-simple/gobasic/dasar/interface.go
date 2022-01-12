package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type count interface {
	area() float64
	around() float64
}

type circle struct {
	diameter float64
}

func (c circle) circleRadius() float64 {
	return c.diameter / 2
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.circleRadius(), 2)
}

func (c circle) around() float64 {
	return math.Pi * c.diameter
}

type square struct {
	side float64
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (s square) around() float64 {
	return s.side * 4
}

type count2d interface {
	area() float64
	around() float64
}

type count3d interface {
	volume() float64
}

type counts interface {
	count2d
	count3d
}

type cube struct {
	side float64
}

func (cb *cube) volume() float64 {
	return math.Pow(cb.side, 3)
}

func (cb *cube) area() float64 {
	return math.Pow(cb.side, 2) * 6
}

func (cb *cube) around() float64 {
	return cb.side * 12
}

type person2 struct {
	name string
	age  int
}

type student2 struct {
	Name  string
	Grade int
}

func main() {
	var figure count

	figure = square{5.0}
	fmt.Println(figure.area(), figure.around())
	figure = circle{14.0}
	var figureCircle circle = figure.(circle)
	fmt.Println(figure.area(), figure.around(), figure.(circle).circleRadius(), figureCircle.circleRadius())

	var figures counts = &cube{2}
	fmt.Println(figures.area(), figures.around(), figures.volume())

	var secret interface{}
	secret = "Setia"
	fmt.Println(secret)
	secret = []string{"Apple", "Orange"}
	fmt.Println(secret)
	secret = 12
	fmt.Println(secret)

	var data map[string]interface{}
	data = map[string]interface{}{
		"name":    "Ajung",
		"grade":   4,
		"hoobies": []string{"swim", "run"},
	}
	fmt.Println(data)

	var number = secret.(int) * 5
	fmt.Println(secret, number)
	secret = []string{"Mango", "Banana"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, " is fruit")

	var secrets interface{} = &person2{name: "Setia", age: 20}
	var name = secrets.(*person2).name
	fmt.Println(name)

	var persons = []map[string]interface{}{
		{"name": "Ajung", "age": 19},
		{"name": "Setia", "age": 20},
	}
	for _, each := range persons {
		fmt.Println(each["name"], each["age"])
	}
	var fruits2 = []interface{}{
		map[string]interface{}{"name": "Orange", "total": 5},
		[]string{"Manggo", "Papaya"},
		"Grape",
	}
	for _, each := range fruits2 {
		fmt.Println(each)
	}

	var number2 = 20
	var reflectValue = reflect.ValueOf(number2)
	fmt.Println(reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println(reflectValue.Int())
	}
	var mark = reflectValue.Interface().(int)
	fmt.Println(reflectValue.Type(), reflectValue.Interface(), mark)

	var s1 = &student2{Name: "Ajung", Grade: 4}
	s1.getProperty()
	var reflectValues = reflect.ValueOf(s1)
	var method = reflectValues.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Setia"),
	})
	fmt.Println("name: ", s1.Name)
}

func (s *student2) getProperty() {
	var reflectValue = reflect.ValueOf(s)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()
	for i := 0; i < reflectType.NumField(); i++ {
		fmt.Println(reflectType.Field(i).Name, reflectType.Field(i).Type, reflectValue.Field(i).Interface())
	}
}

func (s *student2) SetName(name string) {
	s.Name = name
}
