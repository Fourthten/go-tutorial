package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"Ajung", 20},
		{"Setia", 21},
		{"Agung", 23},
		{"Anton", 23},
	}

	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("Server started on: http://localhost:9000")
	http.ListenAndServe(":9000", nil)
}
