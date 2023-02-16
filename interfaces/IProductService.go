package interfaces

import (
	"context"
	"github.com/mhaqiw/product-service/domain"
)

type ProductService interface {
	GetAll(ctx context.Context, sort string) (response domain.ProductsResponsePayload, err error)
	AddProduct(ctx context.Context, request domain.ProductRequestPayload) (product domain.Product, err error)
}
