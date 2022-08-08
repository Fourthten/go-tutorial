package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Apple Mac Book",
		ImageURL: "http://google.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Mac Book","image_url":"http://google.com"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
