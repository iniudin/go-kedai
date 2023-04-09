package users

import (
	"context"
	"gorm.io/gorm"
	"katalog/internal/database"
	"katalog/internal/models"
)

type ServiceImpl struct {
	db *gorm.DB
}

func (service *ServiceImpl) FindAll(ctx context.Context, page int, size int) ([]UserResponse, error) {
	var users []models.User
	tx := service.db.WithContext(ctx).Scopes(database.Paginate(page, size)).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return ToUsersResponse(users), nil
}

func (service *ServiceImpl) FindById(ctx context.Context, id uint) (*UserResponse, error) {
	var user models.User
	tx := service.db.WithContext(ctx).First(&user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return ToUserResponse(user), nil
}

func (service *ServiceImpl) Update(ctx context.Context, request UserUpdateRequest) (*UserResponse, error) {
	user := models.User{}
	tx := service.db.WithContext(ctx).Where("id = ?", request.Id).Updates(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return ToUserResponse(user), nil
}

func (service *ServiceImpl) Delete(ctx context.Context, id uint) error {
	var user models.User
	return service.db.WithContext(ctx).Delete(&user, id).Error
}

func NewServiceImpl(db *gorm.DB) *ServiceImpl {
	return &ServiceImpl{db: db}
}
