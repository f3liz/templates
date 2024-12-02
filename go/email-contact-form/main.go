// Still trying to figure this part

package main

import (
    "log"
    "net/http"
    "email-contact-form/handler"  // Handler package
)

func main() {
    http.HandleFunc("/", handler.ContactFormHandler) 
    log.Fatal(http.ListenAndServe(":8080", nil))     
}
