package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	errInvalidToken     = errors.New("invalid token. claims value is invalid")
	errTokenExpired     = errors.New("token expired")
	errInvalidAuth      = errors.New("invalid authorization header")
	AuthorizationHeader = "Authorization"
	BearerToken         = "Bearer"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID     uuid.UUID
	RoleID     uuid.UUID
	ExpireDate int64
}

// ExtractTokenMetadata handler_func to extract metadata from JWT.
func ExtractTokenMetadata(ctx *gin.Context) (*TokenMetadata, error) {
	token, err := verifyToken(ctx)
	if err != nil {
		return nil, err
	}
	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, userIDExist := claims[payloadUserID]
		if !userIDExist {
			return nil, errInvalidToken
		}
		userIdUUID, err := uuid.Parse(userID.(string))
		if err != nil {
			return nil, err
		}
		roleID, roleIdExist := claims[payloadRoleID]
		if !roleIdExist {
			return nil, errInvalidToken
		}
		roleIdUUID, err := uuid.Parse(roleID.(string))
		if err != nil {
			return nil, err
		}
		expireDate, expireDateExist := claims[payloadExpireDate]
		if !expireDateExist {
			return nil, errInvalidToken
		}
		expireDateInt64 := int64(expireDate.(float64))
		if expireDateInt64 < time.Now().Unix() {
			return nil, errTokenExpired
		}
		// User credentials.
		return &TokenMetadata{
			UserID:     userIdUUID,
			ExpireDate: expireDateInt64,
			RoleID:     roleIdUUID,
		}, nil
	}
	return nil, err
}
func ExtractRefreshTokenMetadata(ctx *gin.Context) (*TokenMetadata, error) {
	token, err := verifyToken(ctx)
	if err != nil {
		return nil, err
	}
	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID, userIDExist := claims[payloadUserID]
		if !userIDExist {
			return nil, errInvalidAuth
		}
		// Expires time.
		expires, expiresExist := claims[payloadExpireDate]
		if !expiresExist {
			return nil, errInvalidAuth
		}
		expiresInt64 := int64(expires.(float64))
		if expiresInt64 < time.Now().Unix() {
			return nil, errTokenExpired
		}
		// User credentials.
		return &TokenMetadata{
			UserID:     userID.(uuid.UUID),
			ExpireDate: expiresInt64,
		}, nil
	}
	return nil, err
}

func extractToken(ctx *gin.Context) string {
	bearToken := ctx.GetHeader(AuthorizationHeader)
	token := fmt.Sprintf("%v", bearToken)
	onlyToken := strings.Split(token, " ")
	if len(onlyToken) != 2 {
		return errInvalidAuth.Error()
	}
	if onlyToken[0] != BearerToken {
		return errInvalidAuth.Error()
	}
	return onlyToken[1]
}

func verifyToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := extractToken(ctx)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}
func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(conf.JWT.JWTSecretKey), nil
}
