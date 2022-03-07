package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var urlString = "http://fourthten.org:80/hello?name=setia&age=20"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Printf("url: %s\n", urlString)
	fmt.Printf("protocol: %s\n", u.Scheme)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)
	var name = u.Query()["name"][0]
	var age = u.Query()["age"][0]
	fmt.Printf("name: %s, age: %s\n", name, age)

	var jsonString = `{"Name": "setia", "Age": 20}`
	var jsonData = []byte(jsonString)
	var data User
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("user: ", data.FullName)
	fmt.Println("age: ", data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)
	fmt.Println("user: ", data1["Name"])
	fmt.Println("age: ", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)
	var decodedData = data2.(map[string]interface{})
	fmt.Println("user: ", decodedData["Name"])
	fmt.Println("age: ", decodedData["Age"])

	var jsonString1 = `[
		{"Name": "Setia", "Age": 20},
		{"Name": "Ajung", "Age": 21}
	]`
	var data3 []User
	var err1 = json.Unmarshal([]byte(jsonString1), &data3)
	if err1 != nil {
		fmt.Println(err1.Error())
		return
	}
	fmt.Println("user 1:", data3[0].FullName)
	fmt.Println("user 2:", data3[1].FullName)

	var object = []User{{"Setia", 20}, {"Ajung", 21}}
	var jsonData2, err2 = json.Marshal(object)
	if err2 != nil {
		fmt.Println(err2.Error())
		return
	}
	var jsonString2 = string(jsonData2)
	fmt.Println(jsonString2)
}
