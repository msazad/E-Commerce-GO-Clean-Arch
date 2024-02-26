package interfaces

import (
	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
)

type OfferUsecase interface {
	AddNewOffer(model models.CreateOffer) error
	MakeOfferExpire(catID int) error
	GetOffers(page, limit int) ([]domain.Offer, error)
}
