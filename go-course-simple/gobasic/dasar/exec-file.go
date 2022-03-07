package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

var path = "/Users/fourthten/Downloads/test.txt"

func main() {
	// exec
	var output1, _ = exec.Command("ls").Output()
	fmt.Printf(" -> ls\n%s\n", string(output1))
	var output2, _ = exec.Command("pwd").Output()
	fmt.Printf(" -> pwd\n%s\n", string(output2))
	var output3, _ = exec.Command("git", "config", "user.name").Output()
	fmt.Printf(" -> git config user.name\n%s\n", string(output3))

	var output []byte
	if runtime.GOOS == "windows" {
		output, _ = exec.Command("cmd", "/C", "git config user.name").Output()
		fmt.Printf(" -> git config user.name\n%s\n", string(output))
	} else {
		output, _ = exec.Command("bash", "-c", "git config user.name").Output()
		fmt.Printf(" -> git config user.name\n%s\n", string(output))
	}

	// file
	createFile()
	writeFile()
	readFile()
	deleteFile()
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("==> file created", path)
}

func writeFile() {
	// open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()
	// write file
	_, err = file.WriteString("hello\n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("let's try\n")
	if isError(err) {
		return
	}
	// save
	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println("==> file saved")
}

func readFile() {
	// open file
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()
	// read file
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}
	fmt.Println("==> file readed")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("==> file deleted")
}
