package user

import (
	"time"
)

type User struct {
	ID        uint
	Name      string
	Phone     string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUpdateRequest struct {
	Id       uint   `validate:"required"`
	Name     string `validate:"required,min=3,max=32"`
	Phone    string `validate:"required,min=10,max=12"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8,max=25"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToUserResponse(u User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Phone:     u.Phone,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToUsersResponse(usersModel []User) (users []UserResponse) {
	for _, user := range usersModel {
		users = append(users, UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Phone:     user.Phone,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return users
}
