package storage

import (
	"context"
	"task/api/models"
)

type StorageI interface {
	User() UserRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *models.CreateUser) (*models.User, error)
	GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error)
	GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error)
	Update(ctx context.Context, req *models.UpdateUser) (int64, error)
	Delete(ctx context.Context, req *models.UserPrimaryKey) error
	// GetFields(ctx context.Context, req *models.GetListUserRequest) (resp *models.GetField, err error)
}
