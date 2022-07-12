package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"
)

type GetByCategoryUseCase interface {
	FindByCategory(category string) ([]model.Product, error)
}

type getByCategoryUseCase struct {
	repo repository.ProductRepository
}

func (p *getByCategoryUseCase) FindByCategory(category string) ([]model.Product, error) {
	return p.repo.GetByCategory(category)
}

func NewCategoryUseCase(repo repository.ProductRepository) GetByCategoryUseCase {
	return &getByCategoryUseCase{repo:repo}
}
