package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Setia","MiddleName":"Ajung","LastName":"Agung","Age":19,"Married":true}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
