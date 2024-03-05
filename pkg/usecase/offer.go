package usecase

import (
	"errors"

	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/repository/interfaces"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
)


type offerUsecase struct{
	offerRepo interfaces.OfferRepository
}

//constructor function


func NewOfferUsecase(offerRepo interfaces.OfferRepository)services.offerUsecase{
	return &offerUsecase{
		offerRepo: offerRepo,
	}
}

func (offU *offerUsecase)AddNewOffer(model models.CreateOffer)error{
	if err:=offU.offerRepo.AddNewOffer(model);err!=nil{
		return errors.New("adding offer failed")
	}
	return nil
}

func (offU *offerUsecase)MakeOfferExpire(catID int)error{
	if err:=offU.offerRepo.MakeOfferExpired(catID);err!=nil{
		return err
	}
	return nil
}


func (offU *offerUsecase)GetOffers(page,limit int)([]domain.Offer,error){
	offers,err:=offU.offerRepo.GetOffers(page,limit)
	if err!=nil{
		return []domain.Offer{},err
	}
	return offers,nil
}