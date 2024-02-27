package usecase

import (
	"errors"
	"fmt"

	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/msazad/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	adminRepository interfaces.AdminRepository
}

// constructor function
func NewAdminUsecase(adRepo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUsecase{
		adminRepository: adRepo,
	}

}

func (au *adminUsecase)LoginHandler(adminDetails models.AdminLogin)(domain.AdminToken,error){
	//Getting details of the admin based on the email provided

	adminCompareDetails,_:=au.adminRepository.LoginHandler(adminDetails)

	if adminCompareDetails.UserName!=adminDetails.Email{
		return domain.AdminToken{},errors.New("admin not exist")
	}
	//Compare password from database that provided by admin
	err:=bcrypt.CompareHashAndPassword([]byte(adminCompareDetails.Password),[]byte(adminDetails.Password))
	if err!=nil{
		return domain.AdminToken{},err
	}
	var adminDetailsResponse models.AdminDetailsResponse

	// err = copier.Copy(&adminDetailsResponse, &adminCompareDetails)
	// if err != nil {
	// 	return domain.AdminToken{}, err
	// }

	adminDetailsResponse.ID=adminCompareDetails.ID
	adminDetailsResponse.Name=adminCompareDetails.Name
	adminDetailsResponse.Email=adminCompareDetails.UserName

	fmt.Println("admindetails response id",adminDetailsResponse.ID)
	fmt.Println("admindetails response name",adminDetailsResponse.Name)
	fmt.Println("admin details response email",adminDetailsResponse.Email)

	fmt.Println("reached here--------")

	token,refresh,err:=helper.GenerateAdminToken(adminDetailsResponse)
	if err!=nil{
		return domain.AdminToken{},err
	}
	return domain.AdminToken{
		Admin: adminDetailsResponse,
		Token: token,
		RefreshToken: refresh,
	},nil
}
