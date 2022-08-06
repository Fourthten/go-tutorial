package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpAjung NoKTP = "18741982741897419874"
	var marriedStatus Married = true
	fmt.Println(noKtpAjung)
	fmt.Println(marriedStatus)
}
