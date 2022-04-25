// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/machester4/products-leaderboard/internal/application/dto"
	mock "github.com/stretchr/testify/mock"
)

// ProductLeadboard is an autogenerated mock type for the ProductLeadboard type
type ProductLeadboard struct {
	mock.Mock
}

// GetTopProducts provides a mock function with given fields: ctx, limit
func (_m *ProductLeadboard) GetTopProducts(ctx context.Context, limit int) ([]dto.ProductLeadboard, error) {
	ret := _m.Called(ctx, limit)

	var r0 []dto.ProductLeadboard
	if rf, ok := ret.Get(0).(func(context.Context, int) []dto.ProductLeadboard); ok {
		r0 = rf(ctx, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.ProductLeadboard)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementScore provides a mock function with given fields: ctx, productID
func (_m *ProductLeadboard) IncrementScore(ctx context.Context, productID string) error {
	ret := _m.Called(ctx, productID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, productID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}