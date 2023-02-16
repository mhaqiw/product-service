package service

import (
	"context"
	"github.com/mhaqiw/product-service/domain"
	"github.com/mhaqiw/product-service/interfaces"
	"time"
)

type productService struct {
	productRepo    interfaces.ProductRepository
	contextTimeout time.Duration
}

func NewProductService(c interfaces.ProductRepository, timeout time.Duration) interfaces.ProductService {
	return &productService{
		productRepo:    c,
		contextTimeout: timeout,
	}
}

func (p *productService) GetAll(ctx context.Context, sort string) (response domain.ProductsResponsePayload, err error) {

	products, err := p.productRepo.Get(ctx, sort)
	response = domain.ProductsResponsePayload{
		List: products,
	}
	return
}

func (p *productService) AddProduct(ctx context.Context, request domain.ProductRequestPayload) (product domain.Product, err error) {
	isAlreadyExists, err := p.productRepo.CheckIsExistByName(ctx, request.Name)
	if err != nil {
		return
	}
	if isAlreadyExists {
		return product, domain.ErrProductAlreadyExists
	}

	product = domain.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
	}

	err = p.productRepo.Insert(ctx, &product)
	return
}
