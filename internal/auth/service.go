package auth

import (
	"context"
	"gorm.io/gorm"

	"katalog/internal/pkg/jwt"
	"katalog/internal/pkg/password"
)

type ServiceImpl struct {
	db *gorm.DB
}

func (service *ServiceImpl) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	var user models.User
	tx := service.db.WithContext(ctx).Where("email = ?", request.Email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := password.ComparePassword(user.Password, request.Password); err != nil {
		return nil, err
	}

	accessToken, err := jwt.GenerateToken(&user, 1)
	if err != nil {
		return nil, err
	}
	refreshToken, err := jwt.GenerateToken(&user, 90)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func (service *ServiceImpl) Register(ctx context.Context, request RegisterRequest) (*RegisterResponse, error) {
	user := models.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Email:    request.Email,
		Password: password.HashPassword(request.Password),
	}
	tx := service.db.WithContext(ctx).Create(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &RegisterResponse{
		ID:        user.ID,
		Name:      user.Name,
		Phone:     user.Phone,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.CreatedAt,
	}, nil
}

func (service *ServiceImpl) Refresh(ctx context.Context, request AccessTokenRequest) (*AccessTokenResponse, error) {
	var user models.User
	tx := service.db.WithContext(ctx).Where("email = ?", request.Email).First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	accessToken, err := jwt.GenerateToken(&user, 1)
	if err != nil {
		return nil, err
	}
	return &AccessTokenResponse{
		TokenType:   "Bearer",
		AccessToken: accessToken,
	}, nil
}

func NewAuthServiceImpl(db *gorm.DB) *ServiceImpl {
	return &ServiceImpl{db: db}
}
