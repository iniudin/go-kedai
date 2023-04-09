package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"katalog/internal/models"
	"os"
	"time"
)

type Token struct {
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}

type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.User, days int) (*Token, error) {
	readFile, err := os.ReadFile("cert/private.pem")
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(readFile)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	tokenExp := now.AddDate(0, 0, days)
	// Create token with claims
	claims := jwt.NewWithClaims(jwt.SigningMethodRS256, &Claims{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExp),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	})

	// Generate encoded token and send it as response.
	token, err := claims.SignedString(privateKey)
	if err != nil {
		return nil, err
	}

	return &Token{
		Token:     token,
		ExpiredAt: tokenExp,
	}, nil
}
