package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"
)

type PaginationUseCase interface {
	FindLimit(limitNumber *model.FindLimit)([]model.Product, error)
	
}

type paginationUseCase struct {
	repo repository.ProductRepository
}

func (p *paginationUseCase) FindLimit(limitNumber *model.FindLimit)([]model.Product, error){
	return p.repo.Pagination(limitNumber)
}

func NewPaginationUseCase(repo repository.ProductRepository) PaginationUseCase {
	return &paginationUseCase{repo: repo}
}