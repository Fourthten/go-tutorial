package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/jlaffaye/ftp"
)

func getStringEntryType(t ftp.EntryType) string {
	switch t {
	case ftp.EntryTypeFile:
		return "(file)"
	case ftp.EntryTypeFolder:
		return "(folder)"
	case ftp.EntryTypeLink:
		return "(link)"
	default:
		return ""
	}
}

func main() {
	const FTP_ADDR = "127.0.0.1:2121"
	const FTP_USERNAME = "admin"
	const FTP_PASSWORD = "123456"

	// Connect to FTP Server
	conn, err := ftp.Dial(FTP_ADDR)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = conn.Login(FTP_USERNAME, FTP_PASSWORD)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Show All Files On
	fmt.Println("========== PATH ./")

	entries, err := conn.List(".")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, entry := range entries {
		fmt.Println(" ->", entry.Name, getStringEntryType(entry.Type))
	}

	// Change Current Directory to ./file
	fmt.Println("========== PATH ./file")

	err = conn.ChangeDir("./file")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Current Dir is ./file, Show All Files on It
	entries, err = conn.List(".")
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, entry := range entries {
		fmt.Println(" ->", entry.Name, getStringEntryType(entry.Type))
	}

	// Change Current Dir Back to Parent Dir
	err = conn.ChangeDirToParent()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Show File Content
	fmt.Println("========== Show Content of Files")

	fileTest1Path := "test1.txt"
	fileTest1, err := conn.Retr(fileTest1Path)
	if err != nil {
		log.Fatal(err.Error())
	}

	test1ContentInBytes, err := ioutil.ReadAll(fileTest1)
	fileTest1.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(" ->", fileTest1Path, "->", string(test1ContentInBytes))
	fileTest2Path := "file/test3.txt"
	fileTest2, err := conn.Retr(fileTest2Path)
	if err != nil {
		log.Fatal(err.Error())
	}

	test2ContentInBytes, err := ioutil.ReadAll(fileTest2)
	fileTest2.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(" ->", fileTest2Path, "->", string(test2ContentInBytes))

	// Download File
	fmt.Println("========== Download File to Local")

	fileMoviePath := "movie.mp4"
	fileMovie, err := conn.Retr(fileMoviePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	destinationMoviePath := "/Users/fourthten/Downloads/download-movie.mp4"
	f, err := os.Create(destinationMoviePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = io.Copy(f, fileMovie)
	if err != nil {
		log.Fatal(err.Error())
	}
	fileMovie.Close()
	f.Close()

	fmt.Println("File download to", destinationMoviePath)

	// Upload File
	fmt.Println("========== Uplaod File to FTP Server")

	sourceFile := "/Users/fourthten/Downloads/testapp.png"
	f, err = os.Open(sourceFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	destinationFile := "./file/logo.png"
	err = conn.Stor(destinationFile, f)
	if err != nil {
		log.Fatal(err.Error())
	}
	f.Close()

	fmt.Println("File", sourceFile, "uploaded to", destinationFile)
	entries, err = conn.List("./file")
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, entry := range entries {
		fmt.Println(" ->", entry.Name, getStringEntryType(entry.Type))
	}
}

// https://github.com/goftp/server (ftp server)
// go get github.com/goftp/server
// go install github.com/goftp/server/appftpd
// ./appftpd -root /root-path (in $GOPATH/bin)
