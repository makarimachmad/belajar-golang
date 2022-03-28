package main

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "PT. Makarimachmad tutorialabcdefgh@gmail.com"
const CONFIG_AUTH_EMAIL = "makarimachmad@gmail.com"
const CONFIG_AUTH_PASSWORD = "02200880amw"

func main() {
	to := []string{"tutorialabcdefgh@gmail.com", "kontak@makarimachmad.com"}
	cc := []string{"achmad.widyanto@spesoulution.com"}
	subject := "Latihan gomail"
	message := "Halo semua ini hanya uji coba"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Berhasil terkirim")
}

func sendMail(to []string, cc []string, subject, message string) error{
	body := "From: " + CONFIG_SENDER_NAME + "\n" + 
		"To: " + strings.Join(to, ",") + "\n" +
		"CC: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr,  auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}
	return nil
}