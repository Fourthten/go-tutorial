package main

import (
	"log"
	"net/http"
)

func StartNonTLSServer() {
	mux := new(http.ServeMux)
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Redirecting to https://localhost/")
		http.Redirect(w, r, "https://localhost/", http.StatusTemporaryRedirect)
	}))
	http.ListenAndServe(":80", mux)
}

func main() {
	go StartNonTLSServer()
	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, TLS!"))
	})
	log.Println("Server started at :443")
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", mux)
	// err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", mux)
	if err != nil {
		panic(err)
	}
}

// openssl genrsa -out server.key 2048
// openssl ecparam -genkey -name secp384r1 -out server.key
// openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

// Common Name (eg, fully qualified host name) []:localhost
// Email Address []:fourthten@gmail.com

// openssl ecparam -genkey -name secp384r1 -noout -out key.pem
// openssl req -new -x509 -sha256 -key key.pem -out cert.pem -days 3650
