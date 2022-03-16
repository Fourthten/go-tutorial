package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type M map[string]interface{}

func doRequest(url, method string, data interface{}) (interface{}, error) {
	var payload *bytes.Buffer = nil
	if data != nil {
		payload = new(bytes.Buffer)
		err := json.NewEncoder(payload).Encode(data)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	certFile, err := ioutil.ReadFile("./certificate/private.crt")
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(certFile)
	tlsConfig := &tls.Config{RootCAs: caCertPool}
	tlsConfig.BuildNameToCertificate()

	client := new(http.Client)
	client.Transport = &http.Transport{
		TLSClientConfig: tlsConfig,
	}

	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	responseBody := make(M)
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func main() {
	baseURL := "https://localhost:9000"
	method := "POST"
	data := M{"Name": "Setia"}

	responseBody, err := doRequest(baseURL+"/data", method, data)
	if err != nil {
		log.Println("ERROR", err.Error())
		return
	}
	log.Printf("%#v \n", responseBody)
}

// https://gist.github.com/croxton/ebfb5f3ac143cd86542788f972434c96
// openssl genrsa -out private.key 4096
// openssl req -new -sha256 -out private.csr -key private.key -config ssl.conf
// openssl req -text -noout -in private.csr
// openssl x509 -req -sha256 -days 3650 -in private.csr -signkey private.key -out private.crt -extensions req_ext -extfile ssl.conf
