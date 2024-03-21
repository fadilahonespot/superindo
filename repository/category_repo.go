package repository

import (
	"context"

	"github.com/fadilahonespot/superindo/entity"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) (resp []entity.Category, err error)
	FindById(ctx context.Context, id int) (resp entity.Category, err error) 
}

type defaultCategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &defaultCategoryRepo{
		db: db,
	}
}

func (r *defaultCategoryRepo) FindAll(ctx context.Context) (resp []entity.Category, err error) {
	err = r.db.WithContext(ctx).Find(&resp).Error
	return
}

func (s *defaultCategoryRepo) FindById(ctx context.Context, id int) (resp entity.Category, err error) {
	err = s.db.WithContext(ctx).Take(&resp, "id = ?", id).Error
	return
}