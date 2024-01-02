package models

import "github.com/msazad/go-Ecommerce/pkg/domain"

// type AddToCart struct{
// UserID int `json:"user_id"`
// InventoryID int `json:"inventory_id"`
// }
type GetCart struct {
	ProductName     string  `json:"product_name"`
	Category_id     int     `json:"category_id"`
	Quantity        int     `json:"quantity"`
	Total           float64 `json:"total_price"`
	DiscountedPrice float64 `json:"dicounted_price"`
}
type CheckOut struct {
	Address        []domain.Address
	Products       []GetCart
	PaymentMethods []domain.PaymentMethod
	TotalPrice     float64
}
type Order struct {
	AddressID int `json:"address_id"`
	PaymentID int `json:"paymentID"`
}
