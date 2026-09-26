package users

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func GenerateToken(userID uuid.UUID, secret string) (*TokenResponse, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &TokenResponse{
		Token: signedToken,
	}, nil
}
