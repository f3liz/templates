package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

const (
	redactAPIURL = "https://redact.aws.eu.pangea.cloud/v1/redact"
)

var (
	redactToken string
)

type Response struct {
	OK        bool   `json:"ok"`
	Error     string `json:"error"`
	Redacted  string `json:"redacted"`
}

type PageData struct {
	CensoredMessage string
}

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	redactToken = os.Getenv("PANGEA_REDACT_TOKEN")
	if redactToken == "" {
		log.Fatal("PANGEA_REDACT_TOKEN environment variable is missing")
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	// Read and serve the index.html file
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func censorText(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Text string `json:"text"`
	}

	// Parse the JSON request body
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.Text == "" {
		http.Error(w, `{"ok": false, "error": "Missing required field 'text'"}`, http.StatusBadRequest)
		return
	}

	// Call the Pangea Redact API
	client := &http.Client{}
	body := map[string]interface{}{
		"text": requestBody.Text,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, `{"ok": false, "error": "Failed to encode request"}`, http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", redactAPIURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, `{"ok": false, "error": "Failed to create request"}`, http.StatusInternalServerError)
		return
	}

	req.Header.Set("Authorization", "Bearer "+redactToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, `{"ok": false, "error": "Failed to call Redact API"}`, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read response from Redact API
	bodyBytes, err := io.ReadAll(resp.Body)  // Using io.ReadAll instead of ioutil.ReadAll
	if err != nil {
		http.Error(w, `{"ok": false, "error": "Failed to read Redact API response"}`, http.StatusInternalServerError)
		return
	}

	var response Response
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		http.Error(w, `{"ok": false, "error": "Failed to parse Redact API response"}`, http.StatusInternalServerError)
		return
	}

	// Send the redacted text back to the client
	if !response.OK || response.Error != "" {
		http.Error(w, `{"ok": false, "error": "`+response.Error+`"}`, http.StatusBadRequest)
		return
	}

	pageData := PageData{CensoredMessage: response.Redacted}
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, pageData)
}

func main() {
	// Serve the index.html page on GET requests
	http.HandleFunc("/", serveIndex)

	// Handle the censor text POST requests
	http.HandleFunc("/censor", censorText)

	// Start the web server
	port := "8080"
	log.Printf("Server is running on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
