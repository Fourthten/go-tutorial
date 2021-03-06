package app

import (
	"encoding/json"
	// "encoding/xml"

	// "fmt"
	"net/http"

	"go-rest/banking/service"

	"github.com/gorilla/mux"
)

// type Customer struct {
// 	Name    string `json:"full_name"	xml:"name"`
// 	City    string `json:"city"			xml:"city"`
// 	Zipcode string `json:"zip_code"		xml:"zipcode"`
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello World!!")
// }

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// 	customers := []Customer{
	// 		{"Agung", "Bekasi", "17511"},
	// 		{"Setia", "Klaten", "17541"},
	// 	}

	// customers, _ := ch.service.GetAllCustomer()
	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprintf(w, err.Error())

		// w.WriteHeader(err.Code)
		// fmt.Fprintf(w, err.Message)

		// w.Header().Add("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(err.AsMessage())

		writeResponse(w, err.Code, err.AsMessage())
	} else {
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// json.NewEncoder(w).Encode(customer)

		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
