package controller

import (
	"github.com/labstack/echo"
	"github.com/mhaqiw/product-service/domain"
	"github.com/mhaqiw/product-service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestProductHandler_GetAll(t *testing.T) {
	e := echo.New()
	tests := []struct {
		name               string
		mockProductService mocks.ProductService
		statusCodeExpected int
		wantErr            bool
		responseBody       string
	}{
		{
			name: "success",
			mockProductService: func() mocks.ProductService {
				service := new(mocks.ProductService)
				date, _ := time.Parse("2006-01-02", "2021-11-22")
				service.On("GetAll", mock.Anything, mock.Anything).Return(domain.ProductsResponsePayload{
					List: []domain.Product{
						{ID: 1, Name: "Iphone", Price: 10.21, Description: "test", CreatedAt: date},
					},
				}, nil)
				return *service
			}(),
			statusCodeExpected: http.StatusOK,
			wantErr:            false,
			responseBody:       "{\"list\":[{\"id\":1,\"name\":\"Iphone\",\"price\":10.21,\"description\":\"test\",\"created_at\":\"2021-11-22T00:00:00Z\"}]}\n",
		},
		{
			name: "error",
			mockProductService: func() mocks.ProductService {
				service := new(mocks.ProductService)
				service.On("GetAll", mock.Anything, mock.Anything).Return(domain.ProductsResponsePayload{
					List: []domain.Product{},
				}, domain.ErrSqlError)
				return *service
			}(),
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            true,
			responseBody:       "{\"message\":\"sql error\"}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewProductHandler(e, &tt.mockProductService)
			req := httptest.NewRequest(http.MethodGet, "/products", nil)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			require.Equal(t, tt.statusCodeExpected, rec.Code)
			require.Equal(t, tt.responseBody, rec.Body.String())
		})
	}
}

func TestProductHandler_Post(t *testing.T) {
	e := echo.New()
	tests := []struct {
		name               string
		mockProductService mocks.ProductService
		statusCodeExpected int
		wantErr            bool
		responseBody       string
		reqBody            io.Reader
	}{
		{
			name: "success",
			mockProductService: func() mocks.ProductService {
				service := new(mocks.ProductService)
				service.On("AddProduct", mock.Anything, mock.Anything).Return(domain.Product{}, nil)
				return *service
			}(),
			statusCodeExpected: http.StatusCreated,
			wantErr:            false,
			reqBody:            strings.NewReader(`{ "name": "Iphone","price": 20.00,"description": "Test" }`),
			responseBody:       "{\"id\":0,\"name\":\"\",\"price\":0,\"description\":\"\",\"created_at\":\"0001-01-01T00:00:00Z\"}\n",
		},
		{
			name: "error sql",
			mockProductService: func() mocks.ProductService {
				service := new(mocks.ProductService)
				service.On("AddProduct", mock.Anything, mock.Anything).Return(domain.Product{}, domain.ErrSqlError)
				return *service
			}(),
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            true,
			reqBody:            strings.NewReader(`{ "name": "Iphone","price": 20.00,"description": "Test" }`),
			responseBody:       "{\"message\":\"sql error\"}\n",
		},
		{
			name: "error invalid body request",
			mockProductService: func() mocks.ProductService {
				service := new(mocks.ProductService)
				service.On("AddProduct", mock.Anything, mock.Anything).Return(domain.Product{}, domain.ErrSqlError)
				return *service
			}(),
			statusCodeExpected: http.StatusUnprocessableEntity,
			wantErr:            true,
			responseBody:       "{\"message\":\"code=400, message=Request body can't be empty\"}\n",
		},
		{
			name: "error empty name",
			mockProductService: func() mocks.ProductService {
				service := new(mocks.ProductService)
				service.On("AddProduct", mock.Anything, mock.Anything).Return(domain.Product{}, domain.ErrSqlError)
				return *service
			}(),
			statusCodeExpected: http.StatusBadRequest,
			wantErr:            true,
			reqBody:            strings.NewReader(`{ "name": "","price": 20.00,"description": "Test" }`),
			responseBody:       "{\"message\":\"Key: 'ProductRequestPayload.Name' Error:Field validation for 'Name' failed on the 'required' tag\"}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewProductHandler(e, &tt.mockProductService)
			req := httptest.NewRequest(http.MethodPost, "/products", tt.reqBody)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)
			require.Equal(t, tt.statusCodeExpected, rec.Code)
			require.Equal(t, tt.responseBody, rec.Body.String())
		})
	}
}
