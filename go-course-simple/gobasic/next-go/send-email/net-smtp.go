package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Company <youremail@gmail.com>"

// const CONFIG_AUTH_EMAIL = "youremail@gmail.com"
// const CONFIG_AUTH_PASSWORD = "yourpass"

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	// auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	// err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	c, err := smtp.Dial(smtpAddr)
	if c != nil {
		defer c.Quit()
	}
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.Mail(CONFIG_SENDER_NAME)
	if err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(body))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	to := []string{"newemail@gmail.com", "otheremail@gmail.com"}
	cc := []string{"anyemail@gmail.com"}
	subject := "Test email"
	message := "This is a test email"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Mail sent")
}
