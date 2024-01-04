package repository

import (
	"fmt"

	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type couponRepository struct{
	db *gorm.DB
}

func NewCouponRepository(DB *gorm.DB)interfaces.CouponRepository{
	return &couponRepository{
		db:DB,
	}
}

func (c *couponRepository)AddCoupon(coup domain.Coupon)error{
	if err:=c.db.Exec("INSERT INTO coupons(name,discount_rate,valid) values($1,$2,$3)",coup.Name,coup.DiscountRate,coup.Valid).Error;err!=nil{
		return err
	}
	return nil
}

func (c *couponRepository)MakeCouponInvalid(id int)error{
	if err:=c.db.Exec("UPDATE coupons SET valid=false where id=$1",id).Error;err!=nil{
		return err
	}

	return nil
}
func (c *couponRepository)FindCouponDiscount(coupon string)int{
	var discountRate int
	err:=c.db.Raw("select discount_rate from coupons where name=?",coupon).Scan(&discountRate).Error
	if err!=nil{
		return 0
	}
	// if !coupon.Valid {
	// return 1
    // }
	fmt.Println("Discount:",discountRate)
	return discountRate
}