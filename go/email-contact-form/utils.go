package main

import (
    "fmt"
    // "log"
    "os"
    "net/smtp"
)

// sendEmail sends an email using SMTP credentials from environment variables
func sendEmail(to, from, subject, body string) error {
    smtpHost := os.Getenv("SMTP_HOST")
    smtpPort := os.Getenv("SMTP_PORT")
    smtpUser := os.Getenv("SMTP_USERNAME")
    smtpPass := os.Getenv("SMTP_PASSWORD")

    msg := []byte("From: " + from + "\r\n" +
        "To: " + to + "\r\n" +
        "Subject: " + subject + "\r\n" +
        "\r\n" + body + "\r\n")

    auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
    err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
    return err
}

// templateFormMessage builds a message from form data
func templateFormMessage(form map[string]string) string {
    var message string
    for key, value := range form {
        message += fmt.Sprintf("%s: %s\n", key, value)
    }
    return message
}