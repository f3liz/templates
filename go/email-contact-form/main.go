package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "strings"
    "github.com/joho/godotenv"
)

func isOriginPermitted(req *http.Request) bool {
    allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
    if allowedOrigins == "" || allowedOrigins == "*" || req.Header.Get("Origin") == "" {
        return true
    }
    allowedOriginsArray := strings.Split(allowedOrigins, ",")
    return contains(allowedOriginsArray, req.Header.Get("Origin"))
}

func contains(arr []string, str string) bool {
    for _, v := range arr {
        if v == str {
            return true
        }
    }
    return false
}

func getCorsHeaders(req *http.Request) map[string]string {
    if req.Header.Get("Origin") == "" {
        return nil
    }
    allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
    origin := req.Header.Get("Origin")
    if allowedOrigins == "" || allowedOrigins == "*" {
        origin = "*"
    }
    return map[string]string{
        "Access-Control-Allow-Origin": origin,
    }
}

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Check environment variables
    requiredEnvVars := []string{"SUBMIT_EMAIL", "SMTP_HOST", "SMTP_USERNAME", "SMTP_PASSWORD"}
    for _, envVar := range requiredEnvVars {
        if os.Getenv(envVar) == "" {
            log.Fatalf("Error: Missing required environment variable: %s", envVar)
        }
    }

    // Handle incoming requests
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        // Check if it's a GET request
        if req.Method == "GET" {
            // Serve the static HTML file
            http.ServeFile(w, req, "./static/index.html")
            return
        }
        
        // If method isn't GET, return a 404 response
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(w, "404 - Page Not Found")
    })

    log.Println("Server started at http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Error starting server:", err)
    }
}
