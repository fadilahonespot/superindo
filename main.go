package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fadilahonespot/superindo/repository"
	"github.com/fadilahonespot/superindo/server/handler"
	"github.com/fadilahonespot/superindo/server/router"
	"github.com/fadilahonespot/superindo/usecase"
	"github.com/fadilahonespot/superindo/utils/database"
	"github.com/fadilahonespot/superindo/utils/logger"
	"github.com/fadilahonespot/superindo/utils/redis"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load Env
	godotenv.Load()

	// Setup database
	db := database.InitDB()

	// Setup logger
	logger.NewLogger()

	// Setup Redis
	redisService := redis.NewCache()

	// Setup repository
	productRepo := repository.NewProductRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)

	// Setup usecase
	productUsecase := usecase.NewProductRepository(productRepo, categoryRepo)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo, redisService)

	// Set handler
	productHandler := handler.NewProductHandler(productUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	// Set Router
	e := echo.New()
	router := router.DefaultRouter{
		ProductHandler:  &productHandler,
		CategoryHandler: categoryHandler,
	}
	router.NewRouter(e).Validate()
	err := e.Start(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatal(err)
	}
}
