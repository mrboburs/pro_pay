package atmost

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pro_pay/model"
	"strings"
)

func GetToken() (out *model.AtmostTokenResponse) {
	// URL to which the POST request will be sent
	url := "https://partner.atmos.uz/token"

	// Create form data
	data := strings.NewReader("grant_type=client_credentials")

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new POST request
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	var base64EncodedString = EncodeCredentialsToBase64()
	// Set the Authorization header
	req.Header.Set("Authorization", "Basic "+base64EncodedString)

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Parse the response body

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Print the access token
	// s.loggers.Info(out)

	return out
}
