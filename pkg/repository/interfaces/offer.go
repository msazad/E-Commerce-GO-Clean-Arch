package interfaces

import (
	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
)

type OfferRepository interface {
	AddNewOffer(offer models.CreateOffer) error
	MakeOfferExpire(catID int) error
	FindDiscountPercentage(catID int) (int, error)
	GetOffers(page, limit int) ([]domain.Offer, error)
}
