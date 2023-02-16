package interfaces

import (
	"context"
	"github.com/mhaqiw/product-service/domain"
)

type ProductRepository interface {
	Get(ctx context.Context, sort string) (product []domain.Product, err error)
	Insert(ctx context.Context, product *domain.Product) (err error)
	CheckIsExistByName(ctx context.Context, productName string) (isAlreadyExist bool, err error)
}
