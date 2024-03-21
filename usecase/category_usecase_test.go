package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/fadilahonespot/superindo/entity"
	"github.com/fadilahonespot/superindo/usecase/dto"
	"github.com/fadilahonespot/superindo/utils/logger"
	"github.com/fadilahonespot/superindo/utils/mocks"
	"github.com/stretchr/testify/mock"
)

func Test_defaultCategoryUsecase_GetCategories(t *testing.T) {
	logger.NewLogger()
	ctx := context.TODO()
	categories := []entity.Category{
		{
			ID:   1,
			Name: "sayuran",
		},
		{
			ID:   2,
			Name: "buahan",
		},
	}

	wantResp := []dto.CategoryResponse{
		{
			Id:   1,
			Name: "sayuran",
		},
		{
			Id:   2,
			Name: "buahan",
		},
	}
	wantRespByte, _ := json.Marshal(wantResp)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name                string
		args                args
		cacheGetResp        string
		cachedGetErr        error
		cachedSetErr        error
		findAllCategoryResp []entity.Category
		findAllCategorErr   error
		wantResp            []dto.CategoryResponse
		wantErr             bool
	}{
		{
			name: "find all category error",
			args: args{
				ctx: ctx,
			},
			cacheGetResp:      "",
			findAllCategorErr: errors.New("error find all category"),
			wantErr:           true,
		},
		{
			name: "success find all category set redis",
			args: args{
				ctx: ctx,
			},
			cacheGetResp:        "",
			findAllCategoryResp: categories,
			wantResp:            wantResp,
			wantErr:             false,
		},
		{
			name: "error unmarshal redis data",
			args: args{
				ctx: ctx,
			},
			cacheGetResp:        "error",
			findAllCategoryResp: categories,
			wantErr:             true,
		},
		{
			name: "success find all category with redis data",
			args: args{
				ctx: ctx,
			},
			cacheGetResp:        string(wantRespByte),
			findAllCategoryResp: categories,
			wantResp:            wantResp,
			wantErr:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cached := new(mocks.CacheService)
			categoryRepo := new(mocks.CategoryRepository)

			cached.On("Get", mock.Anything, mock.Anything).Return(tt.cacheGetResp, tt.cachedGetErr).Once()
			cached.On("Set", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tt.cachedSetErr).Once()
			categoryRepo.On("FindAll", mock.Anything).Return(tt.findAllCategoryResp, tt.findAllCategorErr).Once()

			svc := NewCategoryUsecase(categoryRepo, cached)

			gotResp, err := svc.GetCategories(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("defaultCategoryUsecase.GetCategories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("defaultCategoryUsecase.GetCategories() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
