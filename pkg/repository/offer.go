package repository

import (
	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/repository/interfaces"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type offerRepository struct {
	DB *gorm.DB
}

func NewOfferRepository(DB *gorm.DB) interfaces.OfferRepository {
	return &offerRepository{
		DB: DB,
	}
}

func (o *offerRepository) AddNewOffer(offer models.CreateOffer) error {
	if err := o.DB.Exec("INSERT INTO offers(category_id,discount_rate) values($1,$2)", offer.CategoryID, offer, offer.Discount).Error; err != nil {
		return err
	}
	return nil
}

func (o *offerRepository) MakeOfferExpire(catID int) error {
	if err := o.DB.Exec("UPDATE offers SET valid=false where id=$1", catID).Error; err != nil {
		return err
	}
	return nil
}

func (o *offerRepository) FindDiscountPercentage(catID int) (int, error) {
	var percentage int
	err := o.DB.Raw("select discount_rate from offers where category_id=$1 and valid=true", catID).Scan(&percentage).Error
	if err != nil {
		return 0, err
	}
	return percentage, nil
}

func (o *offerRepository) GetOffers(page, limit int) ([]domain.Offer, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	var offers []domain.Offer

	if err := o.DB.Raw("select id,category_id,discount_rate,valid from offers limit ? offset ?", limit, offset).Scan(&offers).Error; err != nil {
		return []domain.Offer{}, err
	}

	return offers, nil
}
