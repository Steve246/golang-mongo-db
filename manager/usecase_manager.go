package manager

import "golang-mongodb/usecase"

type UseCaseManager interface {
	ProductRegistrationUseCase() usecase.ProductRegistrationUseCase
	//nambain pagination
	PaginationUseCase() usecase.PaginationUseCase
	//update produk
	UpdateProductUseCase() usecase.UpdateProductUsecase

	//delete produk
	DeleteProductUsecase() usecase.DeleteUseCase

	//getById

	GetByIdUseCase() usecase.GetByIdUseCase

	//Getbycategory 

	GetByCategoryUsecase() usecase.GetByCategoryUseCase


}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u useCaseManager) GetByCategoryUsecase() usecase.GetByCategoryUseCase {
	return usecase.NewCategoryUseCase(u.repoManager.ProductRepo())
}

func (u useCaseManager) GetByIdUseCase() usecase.GetByIdUseCase {
	return usecase.NewGetByIdUseCase(u.repoManager.ProductRepo())
}


func (u useCaseManager) DeleteProductUsecase() usecase.DeleteUseCase {
	return usecase.NewDeleteUseCase(u.repoManager.ProductRepo())
}

func(u useCaseManager) UpdateProductUseCase() usecase.UpdateProductUsecase{
	return usecase.NewUpdateProduct(u.repoManager.ProductRepo())
}

func(u useCaseManager) PaginationUseCase() usecase.PaginationUseCase {
	return usecase.NewPaginationUseCase(u.repoManager.ProductRepo())
}


func(u useCaseManager) ProductRegistrationUseCase() usecase.ProductRegistrationUseCase {
	return usecase.NewProductRegistrationUseCase(u.repoManager.ProductRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return useCaseManager{repoManager: repoManager}
}