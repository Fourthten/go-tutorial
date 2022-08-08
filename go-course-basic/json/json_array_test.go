package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Setia",
		MiddleName: "Ajung",
		LastName:   "Agung",
		Hobbies:    []string{"Gaming", "Swimming", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Setia","MiddleName":"Ajung","LastName":"Agung","Age":19,"Married":true,"Hobbies":["Gaming","Swimming","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Setia",
		Addresses: []Address{
			{
				Street:     "Jalan Kosong",
				Country:    "Indonesia",
				PostalCode: "17600",
			},
			{
				Street:     "Jalan Kanan",
				Country:    "Indonesia",
				PostalCode: "18520",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Setia","MiddleName":"Ajung","LastName":"Agung","Age":19,"Married":true,"Hobbies":null,"Addresses":[{"Street":"Jalan Kosong","Country":"Indonesia","PostalCode":"17600"},{"Street":"Jalan Kanan","Country":"Indonesia","PostalCode":"18520"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Kosong","Country":"Indonesia","PostalCode":"17600"},{"Street":"Jalan Kanan","Country":"Indonesia","PostalCode":"18520"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Kosong",
			Country:    "Indonesia",
			PostalCode: "17600",
		},
		{
			Street:     "Jalan Kanan",
			Country:    "Indonesia",
			PostalCode: "18520",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
