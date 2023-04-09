package auth

import (
	"katalog/internal/pkg/jwt"
	"time"
)

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=25"`
}

type RegisterRequest struct {
	Name     string `validate:"required,min=3,max=32"`
	Phone    string `validate:"required,min=10,max=12"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=25"`
}

type LoginResponse struct {
	TokenType    string     `json:"tokenType"`
	RefreshToken *jwt.Token `json:"refreshToken"`
	AccessToken  *jwt.Token `json:"accessToken"`
}

type RegisterResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type AccessTokenRequest struct {
	Id    uint
	Email string
}

type AccessTokenResponse struct {
	TokenType   string     `json:"tokenType"`
	AccessToken *jwt.Token `json:"accessToken"`
}
