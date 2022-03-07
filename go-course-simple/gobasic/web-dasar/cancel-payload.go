package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	done := make(chan bool)
	go func() {
		body, err := ioutil.ReadAll(r.Body)
		_ = err
		_ = body
		time.Sleep(10 * time.Second)
		done <- true
	}()
	select {
	case <-r.Context().Done():
		if err := r.Context().Err(); err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "canceled") {
				log.Println("Request canceled")
			} else {
				log.Println("Unknown error occured.", err.Error())
			}
		}
	case <-done:
		log.Println("done")
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}

// curl -X POST http://localhost:8080/ -H 'Content-Type: application/json' -d '{}'
