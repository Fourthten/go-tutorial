package json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Ajung",
		MiddleName: "Setia",
		LastName:   "Agung",
	}

	encoder.Encode(customer)
}
