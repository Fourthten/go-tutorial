package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var tempPath2 = filepath.Join(os.Getenv("TEMP"), "module-pipeline-temp")

func main() {
	log.Println("start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func proceed() {
	counterTotal := 0
	counterRenamed := 0
	err := filepath.Walk(tempPath2, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// if sub directory, return immediately
		if info.IsDir() {
			return nil
		}
		counterTotal++

		// read file
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// sum it
		sum := fmt.Sprintf("%x", md5.Sum(buf))

		// rename file
		destinationPath := filepath.Join(tempPath2, fmt.Sprintf("file-%s.txt", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}
		counterRenamed++
		return nil
	})
	if err != nil {
		log.Println("ERROR:", err.Error())
	}
	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
}

// go run 2-find-md5-sum-of-file-then-rename-it.go
