package main

import (
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var data = "setia ajung"
	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)
	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString2 = string(encoded)
	fmt.Println(encodedString2)
	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println((err.Error()))
	}
	var decodedString2 = string(decoded)
	fmt.Println(decodedString2)

	var data2 = "https://fourthten.org"
	var encodedString3 = base64.URLEncoding.EncodeToString([]byte(data2))
	fmt.Println(encodedString3)
	var decodedByte2, _ = base64.URLEncoding.DecodeString(encodedString3)
	var decodedString3 = string(decodedByte2)
	fmt.Println(decodedString3)

	// SHA
	var text = "this is secret"
	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	fmt.Println(encryptedString)

	fmt.Printf("original: %s\n\n", text)
	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1: %s\n\n", hashed1)
	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2: %s\n\n", hashed2)
	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3: %s\n\n", hashed3)
	_, _, _ = salt1, salt2, salt3

	// Argument
	// go run file.go arg1 arg2 "arg 3"
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)
	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)

	// Flag
	// go run file.go -name="setia" -age=20
	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int64("age", 25, "type your age")
	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)

	// go run file.go --help
	var data1 = flag.String("name", "anonymous", "type your name")
	fmt.Println(*data1)
	var data3 string
	flag.StringVar(&data3, "gender", "male", "type your gender")
	fmt.Println(data3)
}

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)
	return fmt.Sprintf("%x", encrypted), salt
}

//	Nama Fungsi									Return Value
//	flag.Bool(name, defaultValue, usage)		*bool
//	flag.Duration(name, defaultValue, usage)	*time.Duration
//	flag.Float64(name, defaultValue, usage)		*float64
//	flag.Int(name, defaultValue, usage)			*int
//	flag.Int64(name, defaultValue, usage)		*int64
//	flag.String(name, defaultValue, usage)		*string
//	flag.Uint(name, defaultValue, usage)		*uint
// 	flag.Uint64(name, defaultValue, usage)		*uint64
