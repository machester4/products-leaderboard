package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/machester4/products-leaderboard/internal/application/dto"
	"github.com/machester4/products-leaderboard/internal/application/repositories/mocks"
	"github.com/machester4/products-leaderboard/internal/application/services"
	"github.com/stretchr/testify/mock"
)

// Test product leadboard service

// GetTopProducts when limit is 0 should use default limit of 5
func TestGetTopProductsWhenLimitIs0ShouldUseDefaultLimitOf5(t *testing.T) {
	// Arrange
	limit := 0
	repo := &mocks.ProductLeadboard{}
	service := services.NewProductLeadboard(repo)

	// Act
	repo.On("GetTopProducts", context.Background(), 5).Return([]dto.ProductLeadboard{}, nil)
	service.GetTopProducts(context.Background(), limit)

	// Assert
	repo.AssertExpectations(t)
}

// GetTopProducts when repository returns error should return error
func TestGetTopProductsWhenRepositoryReturnsErrorShouldReturnError(t *testing.T) {
	// Arrange
	limit := 10
	repo := &mocks.ProductLeadboard{}
	service := services.NewProductLeadboard(repo)

	// Act
	repo.On("GetTopProducts", mock.Anything, mock.Anything).Return(nil, errors.New("error"))
	_, err := service.GetTopProducts(context.Background(), limit)

	// Assert
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// GetTopProducts when repository returns valid data should return valid data
func TestGetTopProductsWhenRepositoryReturnsValidDataShouldReturnValidData(t *testing.T) {
	// Arrange
	limit := 10
	repo := &mocks.ProductLeadboard{}
	service := services.NewProductLeadboard(repo)

	// Act
	repo.On("GetTopProducts", mock.Anything, mock.Anything).Return([]dto.ProductLeadboard{
		{
			ProductID: "product-id",
			Score:     1,
		},
	}, nil)
	output, err := service.GetTopProducts(context.Background(), limit)

	// Assert
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	if len(output) != 1 {
		t.Errorf("Expected 1, got %v", len(output))
	}
}

// IncrementScore when repository returns error should return error
func TestIncrementScoreWhenRepositoryReturnsErrorShouldReturnError(t *testing.T) {
	// Arrange
	productID := "product-id"
	repo := &mocks.ProductLeadboard{}
	service := services.NewProductLeadboard(repo)

	// Act
	repo.On("IncrementScore", mock.Anything, mock.Anything).Return(errors.New("error"))
	err := service.IncrementScore(context.Background(), productID)

	// Assert
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

// IncrementScore when repository works correctly should return nil
func TestIncrementScoreWhenRepositoryWorksCorrectlyShouldReturnNil(t *testing.T) {
	// Arrange
	productID := "product-id"
	repo := &mocks.ProductLeadboard{}
	service := services.NewProductLeadboard(repo)

	// Act
	repo.On("IncrementScore", mock.Anything, mock.Anything).Return(nil)
	err := service.IncrementScore(context.Background(), productID)

	// Assert
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
