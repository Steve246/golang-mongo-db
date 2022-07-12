package usecase

import (
	"golang-mongodb/model"
	"golang-mongodb/repository"
)

// UpdateProduct(id string, updateProduct *model.Product) (*model.Product,error)
type UpdateProductUsecase interface {
	Update (id string, updateProduct *model.Product) (*model.Product,error)
}

type updateProductUsecase struct {
	repo repository.ProductRepository
}

func (p *updateProductUsecase) 	Update (id string, updateProduct *model.Product) (*model.Product,error){
	return p.repo.UpdateProduct(id, updateProduct)
}


func NewUpdateProduct(repo repository.ProductRepository) UpdateProductUsecase {
	return &updateProductUsecase{repo:repo}
}