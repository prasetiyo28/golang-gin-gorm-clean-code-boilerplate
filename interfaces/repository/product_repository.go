package repository

import (
	"test1/entities"
)

type ProductRepository struct {
	ProductList []entities.Product
	OrderList   []entities.Order
}

func (r *ProductRepository) CreateProduct(product entities.Product) (entities.Product, error) {
	r.ProductList = append(r.ProductList, product)
	return product, nil
}

func (r *ProductRepository) GetProduct() (*[]entities.Product, error) {
	return &r.ProductList, nil
}

func (r *ProductRepository) CreateOrder(order entities.Order) (entities.Order, error) {

	for i, v := range r.ProductList {
		if order.SKU == v.SKU {
			r.ProductList[i].Qty = v.Qty - order.Qty
		}
	}

	r.OrderList = append(r.OrderList, order)
	return order, nil
}

func (r *ProductRepository) GetOrder() (*[]entities.Order, error) {
	return &r.OrderList, nil
}
