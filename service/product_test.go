package service

import (
	"context"
	"github.com/mhaqiw/product-service/domain"
	"github.com/mhaqiw/product-service/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_productService_Insert(t *testing.T) {
	tests := []struct {
		name            string
		mockProductRepo mocks.ProductRepository
		request         domain.ProductRequestPayload
		wantProduct     domain.Product
		wantErr         bool
	}{
		{
			name: "sucess",
			mockProductRepo: func() mocks.ProductRepository {
				repo := new(mocks.ProductRepository)
				repo.On("CheckIsExistByName", mock.Anything, "Iphone").Return(false, nil).Once()
				repo.On("Insert", mock.Anything, mock.Anything).Return(nil)
				return *repo
			}(),
			request: domain.ProductRequestPayload{
				Name:        "Iphone",
				Price:       10.33,
				Description: "test",
			},
			wantProduct: domain.Product{
				Name:        "Iphone",
				Price:       10.33,
				Description: "test",
			},
			wantErr: false,
		},
		{
			name: "failed already exists",
			mockProductRepo: func() mocks.ProductRepository {
				repo := new(mocks.ProductRepository)
				repo.On("CheckIsExistByName", mock.Anything, "Iphone").Return(true, nil).Once()
				return *repo
			}(),
			request:     domain.ProductRequestPayload{Name: "Iphone", Price: 10.33, Description: "test"},
			wantProduct: domain.Product{},
			wantErr:     true,
		},
		{
			name: "error sql",
			mockProductRepo: func() mocks.ProductRepository {
				repo := new(mocks.ProductRepository)
				repo.On("CheckIsExistByName", mock.Anything, "Iphone").Return(false, domain.ErrSqlError).Once()
				return *repo
			}(),
			request:     domain.ProductRequestPayload{Name: "Iphone", Price: 10.33, Description: "test"},
			wantProduct: domain.Product{},
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewProductService(&tt.mockProductRepo, 5)
			gotProduct, err := s.AddProduct(context.TODO(), tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProduct, tt.wantProduct) {
				t.Errorf("Create() gotProduct = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}

func Test_productService_GetAll(t *testing.T) {
	tests := []struct {
		name            string
		wantProduct     domain.ProductsResponsePayload
		mockProductRepo mocks.ProductRepository
		sort            string
		wantErr         bool
	}{
		{
			name: "success",
			sort: "",
			mockProductRepo: func() mocks.ProductRepository {
				repo := new(mocks.ProductRepository)
				repo.On("Get", mock.Anything, mock.Anything).Return([]domain.Product{{
					Name: "Iphone",
				}}, nil).Once()
				return *repo
			}(),
			wantProduct: domain.ProductsResponsePayload{List: []domain.Product{{
				Name: "Iphone",
			},
			}},
			wantErr: false,
		},
		{
			name: "error",
			sort: "",
			mockProductRepo: func() mocks.ProductRepository {
				repo := new(mocks.ProductRepository)
				repo.On("Get", mock.Anything, mock.Anything).Return(nil, domain.ErrSqlError).Once()
				return *repo
			}(),
			wantProduct: domain.ProductsResponsePayload{},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewProductService(&tt.mockProductRepo, 5)
			gotProduct, err := s.GetAll(context.TODO(), tt.sort)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotProduct, tt.wantProduct) {
				t.Errorf("Create() gotProduct = %v, want %v", gotProduct, tt.wantProduct)
			}
		})
	}
}
