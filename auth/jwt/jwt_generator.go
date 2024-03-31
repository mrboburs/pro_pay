package jwt

import (
	"pro_pay/config"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var conf = config.Config()

const (
	payloadUserID     = "userID"
	payloadRoleID     = "roleID"
	payloadExpireDate = "expireDate"
	EnvDevelopment    = "development"
)

// Tokens struct to describe tokens object.
type Tokens struct {
	Access  string
	Refresh string
}

// GenerateNewTokens handler_func for generate a new Access & Refresh tokens.
func GenerateNewTokens(userID, roleID uuid.UUID, roleTitle string) (*Tokens, error) {
	// Generate JWT Access token.
	accessToken, err := generateNewAccessToken(userID, roleID, roleTitle)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	// Generate JWT Refresh token.
	refreshToken, err := generateNewRefreshToken(userID)
	if err != nil {
		// Return token generation error.
		return nil, err
	}

	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(userID, roleID uuid.UUID, roleTitle string) (string, error) {
	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims[payloadUserID] = userID
	claims[payloadExpireDate] = time.Now().Add(time.Minute * time.Duration(conf.JWT.JWTSecretKeyExpireMinutes)).Unix()
	claims[payloadRoleID] = roleID

	// Set private token credentials:
	// for _, credential := range credentials {
	// 	claims[credential] = true
	// }

	// in local server access token ttl = 10 days
	// if conf.Server.Environment == EnvDevelopment {
	// 	claims[payloadExpireDate] = time.Now().Add(time.Minute * time.Duration(1000*conf.JWT.JWTSecretKeyExpireMinutes)).Unix()
	// } else {
	// 	// in staging server access token ttl = day
	// 	claims[payloadExpireDate] = time.Now().Add(time.Minute * time.Duration(conf.JWTSecretKeyExpireMinutes)).Unix()
	// }
	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(conf.JWT.JWTSecretKey))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

func generateNewRefreshToken(id uuid.UUID) (string, error) {
	// Create a new SHA256 hash.

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims[payloadUserID] = id
	claims[payloadExpireDate] = time.Now().Add(time.Minute * time.Duration(conf.JWT.JWTSecretKeyExpireMinutes)).Unix()

	// Set private token credentials:
	// for _, credential := range credentials {
	// 	claims[credential] = true
	// }

	// in local server access token ttl = 10 days
	// if conf.Environment == EnvDevelopment {
	// 	claims[payloadExpireDate] = time.Now().Add(time.Minute * time.Duration(100*conf.JWTSecretKeyExpireMinutes)).Unix()
	// } else {
	// 	// in staging server access token ttl = day
	// 	claims[payloadExpireDate] = time.Now().Add(time.Minute * time.Duration(conf.JWTSecretKeyExpireMinutes)).Unix()
	// }
	claims[payloadExpireDate] = time.Now().Add(time.Minute * time.Duration(conf.JWT.JWTSecretKeyExpireMinutes)).Unix()
	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(conf.JWT.JWTSecretKey))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}

// ParseRefreshToken handler_func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
