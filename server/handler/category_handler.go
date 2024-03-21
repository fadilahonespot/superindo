package handler

import (
	"net/http"

	"github.com/fadilahonespot/library/response"
	"github.com/fadilahonespot/superindo/usecase"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase}
}

func (h *CategoryHandler) GetAll(c echo.Context) (err error) {
	ctx := c.Request().Context()
	data, err := h.categoryUsecase.GetCategories(ctx)
	if err != nil {
		return
	}

	resp := response.ResponseSuccess(data)
	return c.JSON(http.StatusOK, resp)
}