package dto

import (
	"time"

	"github.com/google/uuid"
)

type ProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Weight      int    `json:"weight"`
	Price       int    `json:"price"`
	CategoryID  int    `json:"categoryId"`
	Image       string `json:"image"`
}

type ProductListResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Weight      int       `json:"weight"`
	Price       int       `json:"price"`
	Image       string    `json:"image"`
	CategoryId  int       `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type DetailProductResponse struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Weight       int    `json:"weight"`
	Price        int    `json:"price"`
	CategoryID   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Image        string `json:"image"`
	CreatedAt   time.Time `json:"createdAt"`
}
