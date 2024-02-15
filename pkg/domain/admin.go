package domain

import "github.com/msazad/go-Ecommerce/pkg/utils/models"

// Admin represents an administrative user in the system
type Admin struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Username string `json:"name" gorm:"validate:required"`
	Email    string `json:"email" gorm:"validate:required"`
	Password string `json:"password" gorm:"validate:required"`
}
type AdminToken struct {
	Admin        models.AdminDetailsResposnse
	Token        string
	RefreshToken string
}
