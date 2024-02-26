package interfaces

import (
	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
)

type CouponUsecase interface {
	Addcoupon(coupon models.Coupon) error
	MakeCouponInvalid(id int) error
	GetCoupons() ([]domain.Coupon, error)
}
