package usecase

import "golang-mongodb/repository"

type DeleteUseCase interface {
	Delete(id string) error
}

type deleteUseCase struct {
	repo repository.ProductRepository
}

func (p *deleteUseCase) Delete(id string) error {
	return p.repo.DeleteProduct(id)
}

func NewDeleteUseCase(repo repository.ProductRepository) DeleteUseCase {
	return &deleteUseCase{repo:repo}
}
