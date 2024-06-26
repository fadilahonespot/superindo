// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/fadilahonespot/superindo/usecase/dto"
	mock "github.com/stretchr/testify/mock"

	paginate "github.com/fadilahonespot/superindo/utils/paginate"
)

// ProductUsecase is an autogenerated mock type for the ProductUsecase type
type ProductUsecase struct {
	mock.Mock
}

// CreateProduct provides a mock function with given fields: ctx, req
func (_m *ProductUsecase) CreateProduct(ctx context.Context, req dto.ProductRequest) (dto.CreateProductResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateProduct")
	}

	var r0 dto.CreateProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.ProductRequest) (dto.CreateProductResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.ProductRequest) dto.CreateProductResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(dto.CreateProductResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.ProductRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteProduct provides a mock function with given fields: ctx, productId
func (_m *ProductUsecase) DeleteProduct(ctx context.Context, productId string) error {
	ret := _m.Called(ctx, productId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, productId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDetailProduct provides a mock function with given fields: ctx, productId
func (_m *ProductUsecase) GetDetailProduct(ctx context.Context, productId string) (dto.DetailProductResponse, error) {
	ret := _m.Called(ctx, productId)

	if len(ret) == 0 {
		panic("no return value specified for GetDetailProduct")
	}

	var r0 dto.DetailProductResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (dto.DetailProductResponse, error)); ok {
		return rf(ctx, productId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) dto.DetailProductResponse); ok {
		r0 = rf(ctx, productId)
	} else {
		r0 = ret.Get(0).(dto.DetailProductResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, productId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListProduct provides a mock function with given fields: ctx, param
func (_m *ProductUsecase) GetListProduct(ctx context.Context, param paginate.Pagination) ([]dto.ProductListResponse, int64, error) {
	ret := _m.Called(ctx, param)

	if len(ret) == 0 {
		panic("no return value specified for GetListProduct")
	}

	var r0 []dto.ProductListResponse
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, paginate.Pagination) ([]dto.ProductListResponse, int64, error)); ok {
		return rf(ctx, param)
	}
	if rf, ok := ret.Get(0).(func(context.Context, paginate.Pagination) []dto.ProductListResponse); ok {
		r0 = rf(ctx, param)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.ProductListResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, paginate.Pagination) int64); ok {
		r1 = rf(ctx, param)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(context.Context, paginate.Pagination) error); ok {
		r2 = rf(ctx, param)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdateProduct provides a mock function with given fields: ctx, productId, req
func (_m *ProductUsecase) UpdateProduct(ctx context.Context, productId string, req dto.ProductRequest) error {
	ret := _m.Called(ctx, productId, req)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProduct")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, dto.ProductRequest) error); ok {
		r0 = rf(ctx, productId, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewProductUsecase creates a new instance of ProductUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProductUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *ProductUsecase {
	mock := &ProductUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
