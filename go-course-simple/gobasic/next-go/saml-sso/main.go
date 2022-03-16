package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crewjam/saml/samlsp"
)

func landingHandler(w http.ResponseWriter, r *http.Request) {
	name := samlsp.AttributeFromContext(r.Context(), "displayName")
	w.Write([]byte(fmt.Sprintf("Welcome %s!", name)))
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
}

func main() {
	sp, err := newSamlMiddleware()
	if err != nil {
		log.Fatal(err.Error())
	}
	http.Handle("/saml/", sp)

	http.Handle("/index", sp.RequireAccount(http.HandlerFunc(landingHandler)))
	http.Handle("/hello", sp.RequireAccount(http.HandlerFunc(helloHandler)))

	portString := fmt.Sprintf(":%d", webserverPort)
	fmt.Println("Listening on port", portString)
	http.ListenAndServe(portString, nil)
}

// openssl req -x509 -newkey rsa:2048 -keyout myservice.key -out myservice.cert -days 365 -nodes -subj "/CN=localhost"
// download metadata [http://localhost:9000/saml/metadata]
// upload metadata [https://samltest.id/upload.php]
