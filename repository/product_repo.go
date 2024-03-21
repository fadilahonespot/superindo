package repository

import (
	"context"

	"github.com/fadilahonespot/superindo/entity"
	"github.com/fadilahonespot/superindo/utils/paginate"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetListProduct(ctx context.Context, param paginate.Pagination) (resp []entity.Product, count int64, err error)
	GetProductById(ctx context.Context, id string) (resp *entity.Product, err error)
	CreateProduct(ctx context.Context, req *entity.Product) (err error)
	UpdateProduct(ctx context.Context, req *entity.Product) (err error)
	DeleteProduct(ctx context.Context, id string) (err error)
}

type defaultProductRepo struct {
	db *gorm.DB
} 

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &defaultProductRepo{db: db}
}

func (s *defaultProductRepo) GetListProduct(ctx context.Context, param paginate.Pagination) (resp []entity.Product, count int64, err error) {
	query := func (db *gorm.DB) *gorm.DB  {
		if param.Search != "" {
			db.Where("name LIKE ? OR id LIKE ?", "%"+param.Search+"%", "%"+param.Search+"%")
		}

		if param.CategoryId != 0 {
			db.Where("category_id = ?", param.CategoryId)
		}

		if param.MaxPrince != 0 {
			db.Where("price >= ? AND price <= ?", param.MinPrice,  param.MaxPrince)
		}

		if param.CreatedAt != "" {
			db.Where("created_at LIKE ?", param.CreatedAt+"%")
		}

		if param.ProductName != "" && param.Search == "" {
			db.Where("name LIKE ?", "%"+param.ProductName+"%")
		}

		return db
	}

	err = s.db.WithContext(ctx).Model(&entity.Product{}).Scopes(query).Count(&count).Error
	if err != nil {
		return
	}

	err = s.db.WithContext(ctx).Scopes(paginate.Paginate(param.Page, param.Limit)).Scopes(query).Find(&resp).Error
	return
}

func (s *defaultProductRepo) GetProductById(ctx context.Context, id string) (resp *entity.Product, err error) {
	err = s.db.WithContext(ctx).Preload("Category").Take(&resp, "id = ?", id).Error
	return
}

func (s *defaultProductRepo) CreateProduct(ctx context.Context, req *entity.Product) (err error) {
	err = s.db.WithContext(ctx).Create(req).Error
	return
}

func (s *defaultProductRepo) UpdateProduct(ctx context.Context, req *entity.Product) (err error) {
	err = s.db.WithContext(ctx).Save(req).Error
	return
}

func (s *defaultProductRepo) DeleteProduct(ctx context.Context, id string) (err error) {
	err = s.db.WithContext(ctx).Delete(&entity.Product{}, "id = ?", id).Error
	return
}