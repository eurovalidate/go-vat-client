package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.eurovalidate.com/v1/vat/NL820646660B01", nil)
	req.Header.Set("X-API-Key", os.Getenv("EUROVALIDATE_API_KEY"))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	fmt.Printf("Status: %s\n", result["status"])
	fmt.Printf("Company: %s\n", result["company_name"])
}
