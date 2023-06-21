package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// Create request for Digital Divide Solution
	req, err := http.NewRequest("GET", "https://www.example.com/digital-divide-solution", nil) 
	if err != nil {
		fmt.Println("Error: failed to create request.")
	}

	// Set custom headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic ew123fa")

	// Create URL query parameters
	params := url.Values{} 
	params.Add("query_field1", "value1")
	params.Add("query_field2", "value2")
	req.URL.RawQuery = params.Encode()

	// Create HTTP client
	client := &http.Client{} 

	// Send request
	resp, err := client.Do(req) 
	if err != nil {
		fmt.Println("Error: failed to send request.")
	}

	// Read response
	defer resp.Body.Close() 
	bodyBytes, err := ioutil.ReadAll(resp.Body) 
	if err != nil {
		fmt.Println("Error: failed to read response.")
	}
	bodyString := string(bodyBytes)

	// Parse response body
	var responseObj map[string]string
	err = json.Unmarshal(bodyBytes, &responseObj) 
	if err != nil {
		fmt.Println("Error: failed to parse response body.")
	}

	// Handle response
	if strings.ToLower(responseObj["status"]) == "success" {
		fmt.Println("Request successful.")
	} else {
		fmt.Println("Request failed.")
	}
}