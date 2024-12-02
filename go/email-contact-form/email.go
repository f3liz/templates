package email

import (
    "fmt"
    "gopkg.in/gomail.v2"
)

func SendEmail(name, emailAddr, message string) error {
    m := gomail.NewMessage()
    m.SetHeader("From", "youremail@example.com")
    m.SetHeader("To", "recipient@example.com")
    m.SetHeader("Subject", "New Contact Form Submission")
    m.SetBody("text/plain", fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", name, emailAddr, message))

    // Send the email
    d := gomail.NewDialer("smtp.example.com", 587, "youremail@example.com", "yourpassword")
    if err := d.DialAndSend(m); err != nil {
        return err
    }

    return nil
}
