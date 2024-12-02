package handler

import (
    "fmt"
    "net/http"
    "email-contact-form/email" // Import the email package
)

func ContactFormHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        // Handle the POST request (form submission)
        name := r.FormValue("name")
        emailAddr := r.FormValue("email")
        message := r.FormValue("message")

        // Send email
        err := email.SendEmail(name, emailAddr, message)
        if err != nil {
            http.Error(w, "Failed to send email", http.StatusInternalServerError)
            return
        }

        // Send a success response
        fmt.Fprintf(w, "Thank you for contacting us, %s!", name)
    } else {
        // Serve the contact form on GET request
        http.ServeFile(w, r, "static/index.html")
    }
}
