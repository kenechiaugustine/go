package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "os"
)

type BankData struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	ShortCode string `json:"shortcode"`
}

func main() {
	// Make HTTP GET request to the API
	url := "https://api.flutterwave.com/v3/banks/NG"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request:", err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error sending HTTP request:", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading HTTP response:", err)
	}

	// Save the API response data to "data.json" file
	err = ioutil.WriteFile("data.json", body, 0644)
	if err != nil {
		log.Fatal("Error writing to data.json file:", err)
	}
	fmt.Println("Data saved to data.json")
}
