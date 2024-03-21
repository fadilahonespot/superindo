package usecase

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/fadilahonespot/library/errors"
	"github.com/fadilahonespot/superindo/repository"
	"github.com/fadilahonespot/superindo/usecase/dto"
	"github.com/fadilahonespot/superindo/utils/logger"
	"github.com/fadilahonespot/superindo/utils/redis"
)

type CategoryUsecase interface {
	GetCategories(ctx context.Context) (resp []dto.Category, err error)
}

type defaultCategoryUsecase struct {
	categoryRepo repository.CategoryRepository
	redisService redis.CacheService
}

func NewCategoryUsecase(categoryRepo repository.CategoryRepository, redisService redis.CacheService) CategoryUsecase {
	return &defaultCategoryUsecase{categoryRepo, redisService}
}

func (s *defaultCategoryUsecase) GetCategories(ctx context.Context) (resp []dto.Category, err error) {
	key := "GetCategories"
	value, _ := s.redisService.Get(ctx, key)
	if value == "" {
		datas, errRes := s.categoryRepo.FindAll(ctx)
		if errRes != nil {
			logger.Error(ctx, "error get categories", errRes)
			err = errors.SetError(http.StatusInternalServerError, "error find all category")
			return
		}

		for i := 0; i < len(datas); i++ {
			resp = append(resp, dto.Category{
				Id:   datas[i].ID,
				Name: datas[i].Name,
			})
		}

		respByte, _ := json.Marshal(resp)
		s.redisService.Set(ctx, key, string(respByte), time.Duration(24) * time.Hour)
		return
	}

	err = json.Unmarshal([]byte(value), &resp)
	if err != nil {
		logger.Error(ctx, "error unmarshal redis data", err.Error())
		err = errors.SetError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	return
}
