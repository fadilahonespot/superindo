package dto

import (
	"time"
)

type ProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Weight      int    `json:"weight" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	CategoryID  int    `json:"categoryId" validate:"required"`
	Image       string `json:"image" validate:"required"`
}

type ProductListResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Weight      int       `json:"weight"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
	CategoryId  int       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type DetailProductResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Weight       int       `json:"weight"`
	Price        int       `json:"price"`
	CategoryID   int       `json:"categoryId"`
	CategoryName string    `json:"categoryName"`
	Image        string    `json:"image"`
	CreatedAt    time.Time `json:"createdAt"`
}

type CreateProductResponse struct {
	ID string `json:"id"`
}
