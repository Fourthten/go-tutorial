package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Company <youremail@gmail.com>"

// const CONFIG_AUTH_EMAIL = "youremail@gmail.com"
// const CONFIG_AUTH_PASSWORD = "yourpass"

func main() {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "newemail@gmail.com", "otheremail@gmail.com")
	mailer.SetAddressHeader("Cc", "anyemail@gmail.com", "Any Text")
	mailer.SetHeader("Subject", "Hello World")
	mailer.SetBody("text/html", "Hello <b>World</b>")
	mailer.Attach("../gzip/testapp.png")

	// dialer := gomail.NewDialer(CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
	dialer := &gomail.Dialer{Host: CONFIG_SMTP_HOST, Port: CONFIG_SMTP_PORT}
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Mail sent")
}
