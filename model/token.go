package model

type Token struct {
	AccessToken  string
	RefreshToken string
}
type AtmostTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}