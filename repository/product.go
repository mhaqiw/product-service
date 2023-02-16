package repository

import (
	"context"
	"database/sql"
	"github.com/mhaqiw/product-service/domain"
	"github.com/mhaqiw/product-service/interfaces"
)

type productRepository struct {
	Conn *sql.DB
}

func NewProductRepository(Conn *sql.DB) interfaces.ProductRepository {
	return &productRepository{Conn}
}

func (p *productRepository) Get(ctx context.Context, sort string) (res []domain.Product, err error) {
	sortQuery := ""
	switch sort {
	case "newest":
		sortQuery = "ORDER BY created_at DESC"
	case "cheapest":
		sortQuery = "ORDER BY price ASC"
	case "expensive":
		sortQuery = "ORDER BY price DESC"
	case "az":
		sortQuery = "ORDER BY name ASC"
	case "za":
		sortQuery = "ORDER BY name DESC"
	default:
	}

	query := `SELECT id, name, price, description, created_at FROM product  ` + sortQuery

	rows, err := p.Conn.QueryContext(ctx, query)

	if err != nil {
		return
	}

	result := make([]domain.Product, 0)
	for rows.Next() {
		t := domain.Product{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Price,
			&t.Description,
			&t.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (p *productRepository) Insert(ctx context.Context, product *domain.Product) (err error) {
	query := `INSERT INTO product( name, price, description ) VALUES( $1, $2, $3) RETURNING id, created_at`
	err = p.Conn.QueryRowContext(ctx, query, product.Name, product.Price, product.Description).Scan(&product.ID, &product.CreatedAt)
	if err != nil {
		return
	}

	return
}

func (p *productRepository) CheckIsExistByName(ctx context.Context, productName string) (isAlreadyExist bool, err error) {
	query := `SELECT EXISTS(SELECT 1 FROM product WHERE name= $1)`
	stmt, err := p.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	row := stmt.QueryRowContext(ctx, productName)
	if err != nil {
		return false, err
	}
	err = row.Scan(
		&isAlreadyExist,
	)
	return
}
