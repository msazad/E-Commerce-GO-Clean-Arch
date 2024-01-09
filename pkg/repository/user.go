package repository

import (
	"github.com/msazad/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)


type userDatabase struct{
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB)interfaces.UserRepository{
	return &userDatabase{DB:DB}
}

func (c *userDatabase)CheckUserAvailability(email string)bool{
	var count int
	query:=`select count(*) from users where email=?`
	if err:=c.DB.Raw(query,email).Scan(&count).Error;err!=nil{
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count>0
}
func (c *userDatabase) UserBlockStatus(email string)(bool,string){
	var permission bool
	err:=c.DB.Raw("select permission from users where email=?",email).Scan(&permission).Error
}