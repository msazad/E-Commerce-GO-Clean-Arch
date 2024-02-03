package interfaces

import (
	"mime/multipart"

	"github.com/msazad/go-Ecommerce/pkg/utils/models"
)

type InventoryUsecase interface {
	AddInventory(inventory models.Inventory,image *multipart.FileHeader)(models.InventoryResponse,error)
	UpdateInventory(invID int,invData models.UpdateInventory)(models.Inventory,error)
	UpdateImage(invID int,image *multipart.FileHeader)(models.Inventory,error)
	DeleteInventory(id string)error
	ShowIndividualProducts(id string)(models.InventoryDetails,error)
}
