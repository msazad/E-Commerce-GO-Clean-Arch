package interfaces

import (
	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUserByID(id string) (domain.User, error)
	UpdateBlockUserByID(user domain.User) error
	GetUsers(page int, limit int) ([]models.UserDetailsAtAdmin, error)
}
