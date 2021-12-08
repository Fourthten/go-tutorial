package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string, 0)

	// get params
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz, ",")

	if len(timezones) <= 1 {
		// get time location
		loc, err := time.LoadLocation(tz)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("invalid timezone %s", tz)))
		} else {
			// show time
			response["current_time"] = time.Now().In(loc).String()
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		for _, tzdb := range timezones {
			// get time location
			loc, err := time.LoadLocation(tzdb)

			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("invalid timezone %s in input", tzdb)))
				return
			}
			// show time
			now := time.Now().In(loc)
			response[tzdb] = now.String()
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
