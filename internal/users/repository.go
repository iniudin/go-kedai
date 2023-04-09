package users

import (
	"context"
	"database/sql"
)

type RepositoryImpl struct {
	db *sql.DB
}

func (r *RepositoryImpl) Save(ctx context.Context, user User) {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryImpl) FindAll(ctx context.Context, page int, size int) ([]UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryImpl) FindById(ctx context.Context, id uint) (*UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryImpl) Update(ctx context.Context, user User) (*UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RepositoryImpl) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}

func NewRepositoryImpl(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}
