package interfaces

import "github.com/msazad/go-Ecommerce/pkg/domain"

type CategoryRepository interface {
	AddCategory(category string) (domain.Category, error)
	CheckCategory(current string) (bool, error)
	UpdateCategory(current, new string) (domain.Category, error)
	DeleteCategory(categoryID string) error
	GetCategories(page, limit int) ([]domain.Category, error)
}
