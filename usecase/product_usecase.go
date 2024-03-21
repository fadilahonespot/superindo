package usecase

import (
	"context"
	"net/http"

	"github.com/fadilahonespot/library/errors"
	"github.com/fadilahonespot/superindo/entity"
	"github.com/fadilahonespot/superindo/repository"
	"github.com/fadilahonespot/superindo/usecase/dto"
	"github.com/fadilahonespot/superindo/utils/logger"
	"github.com/fadilahonespot/superindo/utils/paginate"
)

type ProductUsecase interface {
	CreateProduct(ctx context.Context, req dto.ProductRequest) (resp dto.CreateProductResponse, err error)
	GetListProduct(ctx context.Context, param paginate.Pagination) (resp []dto.ProductListResponse, count int64, err error)
	GetDetailProduct(ctx context.Context, productId string) (resp dto.DetailProductResponse, err error)
	UpdateProduct(ctx context.Context, productId string, req dto.ProductRequest) (err error)
	DeleteProduct(ctx context.Context, productId string) (err error)
}

type defaultProductUsecase struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductRepository(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductUsecase {
	return &defaultProductUsecase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *defaultProductUsecase) CreateProduct(ctx context.Context, req dto.ProductRequest) (resp dto.CreateProductResponse, err error) {
	_, err = s.categoryRepo.FindById(ctx, req.CategoryID)
	if err != nil {
		logger.Error(ctx, "category product is not exist")
		err = errors.SetError(http.StatusNotFound, "product category not found")
		return
	}
	reqProduct := entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Weight:      req.Weight,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
	}

	err = s.productRepo.CreateProduct(ctx, &reqProduct)
	if err != nil {
		logger.Error(ctx, "error creating product", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	resp.ID = reqProduct.ID
	return
}

func (s *defaultProductUsecase) GetListProduct(ctx context.Context, param paginate.Pagination) (resp []dto.ProductListResponse, count int64, err error) {
	data, count, err := s.productRepo.GetListProduct(ctx, param)
	if err != nil {
		logger.Error(ctx, "error getting product list", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	for i := 0; i < len(data); i++ {
		resp = append(resp, dto.ProductListResponse{
			ID:          data[i].ID,
			Name:        data[i].Name,
			Description: data[i].Description,
			Weight:      data[i].Weight,
			Price:       data[i].Price,
			Image:       data[i].Image,
			CategoryId:  data[i].CategoryID,
			CreatedAt:   data[i].CreatedAt,
		})
	}

	return
}

func (s *defaultProductUsecase) GetDetailProduct(ctx context.Context, productId string) (resp dto.DetailProductResponse, err error) {
	data, err := s.productRepo.GetProductById(ctx, productId)
	if err != nil {
		logger.Error(ctx, "error getting product", err.Error())
		err = errors.SetError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	resp = dto.DetailProductResponse{
		Name:         data.Name,
		Description:  data.Description,
		Weight:       data.Weight,
		Price:        data.Price,
		CategoryID:   data.CategoryID,
		CategoryName: data.Category.Name,
		Image:        data.Image,
		CreatedAt:    data.CreatedAt,
	}

	return
}

func (s *defaultProductUsecase) UpdateProduct(ctx context.Context, productId string, req dto.ProductRequest) (err error) {
	productData, err := s.productRepo.GetProductById(ctx, productId)
	if err != nil {
		logger.Error(ctx, "failed to get product: ", err.Error())
		err = errors.SetError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	_, err = s.categoryRepo.FindById(ctx, req.CategoryID)
	if err != nil {
		logger.Error(ctx, "category product is not exist")
		err = errors.SetError(http.StatusNotFound, "product category not found")
		return
	}

	productData.Name = req.Name
	productData.Description = req.Description
	productData.Image = req.Image
	productData.Price = req.Price
	productData.Weight = req.Weight
	productData.CategoryID = req.CategoryID

	err = s.productRepo.UpdateProduct(ctx, productData)
	if err != nil {
		logger.Error(ctx, "failed to update product", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}

func (s *defaultProductUsecase) DeleteProduct(ctx context.Context, productId string) (err error) {
	_, err = s.productRepo.GetProductById(ctx, productId)
	if err != nil {
		logger.Error(ctx, "failed to get product: ", err.Error())
		err = errors.SetError(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	err = s.productRepo.DeleteProduct(ctx, productId)
	if err != nil {
		logger.Error(ctx, "failed to delete product: ", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	return
}
