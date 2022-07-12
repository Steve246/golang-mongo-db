package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"
)

type GetByIdUseCase interface {
	FindById(id string) (*model.Product, error)
}

type getByIdUseCase struct {
	repo repository.ProductRepository
}

func (p *getByIdUseCase) FindById(id string) (*model.Product, error) {
	return p.repo.GetById(id)
}


func NewGetByIdUseCase(repo repository.ProductRepository) GetByIdUseCase {
	return &getByIdUseCase{repo:repo}
}
