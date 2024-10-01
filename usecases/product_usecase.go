package usecases

import "test1/entities"

type ProductRepository interface {
	CreateProduct(entities.Product) (entities.Product, error)
	CreateOrder(entities.Order) (entities.Order, error)
	GetProduct() (*[]entities.Product, error)
	GetOrder() (*[]entities.Order, error)
}

type ProductUsecase struct {
	ProductRepo ProductRepository
}

func (uc *ProductUsecase) CreateProduct(user entities.Product) (entities.Product, error) {
	return uc.ProductRepo.CreateProduct(user)
}

func (uc *ProductUsecase) GetProduct() (*[]entities.Product, error) {
	return uc.ProductRepo.GetProduct()
}

func (uc *ProductUsecase) CreateOrder(order entities.Order) (entities.Order, error) {
	return uc.ProductRepo.CreateOrder(order)
}

func (uc *ProductUsecase) GetOrder() (*[]entities.Order, error) {
	return uc.ProductRepo.GetOrder()
}
