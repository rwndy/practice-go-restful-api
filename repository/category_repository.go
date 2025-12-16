package repository

import (
	"Users/riwandi/Documents/practice/go-restful-api/model/domain"
	"context"
	"database/sql"
)

// CategoryRepository for contract using an interface
// biasakan menggunakan context
type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
