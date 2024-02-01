package interfaces

import "github.com/msazad/go-Ecommerce/pkg/domain"

type CategoryUsecase interface {
	AddCategory(category string)(domain.Category,error)
	UpdateCategory(current,new string)(domain.Category,error)
}