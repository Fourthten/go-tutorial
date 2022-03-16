package main

import (
	"log"
	"net/http"
	// "golang.org/x/net/http2"
)

func main() {
	mux := new(http.ServeMux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, TLS!"))
	})
	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":9000"

	// http2.ConfigureServer(server, nil) // go 1.6 below

	log.Println("Server started at :9000")
	err := server.ListenAndServeTLS("../tls-web/server.crt", "../tls-web/server.key")
	if err != nil {
		panic(err)
	}
}
