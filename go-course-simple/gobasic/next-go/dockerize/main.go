package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env is required")
	}

	instanceID := os.Getenv("INSTANCE_ID")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "http method not allowed", http.StatusBadRequest)
			return
		}

		text := "Hello world"
		if instanceID != "" {
			text = text + ". from " + instanceID
		}

		w.Write([]byte(text))
	})
	server := new(http.Server)
	server.Handler = mux
	server.Addr = "0.0.0.0:" + port

	log.Println("server starting at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Windows
// set PORT=8080
// set INSTANCE_ID="my first instance"
// go run main.go

// non-Windows
// export PORT=8080
// export INSTANCE_ID="my first instance"
// go run main.go

// https://docs.docker.com/get-docker/

// Usage
// # build image
// docker build -t my-image-hello-world .
// # create container from above image
// docker container create --name my-container-hello-world -e PORT=8080 -e INSTANCE_ID="my first instance" -p 8080:8080 my-image-hello-world
// # run newly created container
// docker container start my-container-hello-world
// # stop container
// docker container stop my-container-hello-world
// # remove container
// docker container rm my-container-hello-world
// # remove image
// docker image rm my-image-hello-world

// Or
// # create container then run it, then auto delete if stopped
// docker container run --name my-container-hello-world --rm -itd -e PORT=8080 -e INSTANCE_ID="my first instance" -p 8080:8080 -d my-image-hello-world

// docker images (show image)
// docker container ls -a (show container)
