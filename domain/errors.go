package domain

import "errors"

var (
	ErrProductAlreadyExists = errors.New("product with that name already exists")
	ErrSqlError             = errors.New("sql error") // for mock error repository
)
