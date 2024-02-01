package interfaces

import "github.com/msazad/go-Ecommerce/pkg/utils/models"


type AdminUsecase interface{
	LoginHandler(admindetails models.AdminLogin)(models.TokenAdmin,error)
	BlockUser(id string)error
	UnblockUser(id string)error
	GetUsers(page,limit int)([]models.UserDetailsAtAdmin,error)
}