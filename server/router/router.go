package router

import (
	"github.com/fadilahonespot/superindo/server/handler"
	"github.com/fadilahonespot/superindo/server/middleware"
	"github.com/labstack/echo/v4"
)

type DefaultRouter struct {
	ProductHandler  *handler.ProductHandler
	CategoryHandler *handler.CategoryHandler
}

func (d *DefaultRouter) Validate() {
	if d.ProductHandler == nil {
		panic("product handler is nil")
	}
	if d.CategoryHandler == nil {
		panic("category handler is nil")
	}
}

func (d *DefaultRouter) NewRouter(e *echo.Echo) *DefaultRouter {
	middleware.SetupMiddleware(e)

	e.POST("/product", d.ProductHandler.AddProduct)
	e.GET("/product", d.ProductHandler.GetListProduct)
	e.GET("/product/:productId", d.ProductHandler.GetProductDetail)
	e.PUT("/product/:productId", d.ProductHandler.UpdateProduct)
	e.DELETE("/product/:productId", d.ProductHandler.DeleteProduct)

	e.GET("/category", d.CategoryHandler.GetAll)

	return d
}
