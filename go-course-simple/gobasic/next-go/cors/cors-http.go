package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}
		w.Write([]byte("hello"))
	})
	log.Println("Starting at :9000")
	http.ListenAndServe(":9000", nil)
}

// Jquery Injector
// $.ajax({ url: "http://localhost:9000/index" })

// let jq = document.createElement("script");
// jq.src = "https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js";
// jq.onload = function() { jQuery.ajax({ url: "http://localhost:9000/index" }) };
// document.body.appendChild(jq);

// PUT, DELETE, CONNECT, OPTIONS, TRACE, PATCH
// Authorization, X-CSRF-Token
// applocation/x-www-form-urlencoded, multipart/form-data, text/plain
