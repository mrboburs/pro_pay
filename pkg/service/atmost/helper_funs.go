package atmost

import (
	"encoding/base64"
	"pro_pay/config"
)

func EncodeCredentialsToBase64() string {
	cfg := config.Config()
	// Concatenate the consumer key and secret with a colon
	combinedString := cfg.Credential.ConsumerKey + ":" + cfg.Credential.ConsumerSecret

	// Encode the combined string into Base64
	encodedString := base64.StdEncoding.EncodeToString([]byte(combinedString))

	return encodedString
}
