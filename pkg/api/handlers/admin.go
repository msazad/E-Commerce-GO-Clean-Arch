package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
	"github.com/msazad/go-Ecommerce/pkg/utils/response"
)

type AdminHandler struct {
	adminUsecase services.AdminUsecase
}

//constructor function

func NewAdminHandler(adminUsecase services.adminUsecase) *adminHandler {
	return &AdminHandler{
		adminUsecase: adminUsecase,
	}
}

// @Summary Adminu Login
// @Description Login handler for admins
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body models.AdminLogin true "Admin login details"
// @Success 200 {object} response.Response{} 
// @Failure 500 {object} response.Response{}
// @Router /admin/adminlogin [post]
func (ah *AdminHandler)LoginHandler(c *gin.Context){
	// login handler for the admin
	var adminDetails models.AdminLogin
	if err:=c.BindJSON(&adminDetails);err!=nil{
		errRes:=response.ClientResponse(http.StatusBadRequest,"details not in correct format",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	//c.Set("",admin.Token)
	//c.Set("Refresh",admin.RefreshToken)
	c.SetCookie("Authorization",admin.Token,3600 *24*30,"/",false,true)
	c.SetCookie("Refresh",admin.RefreshToken,3600 *24*30,"/","",false,true)

	successRes:=response.ClientResponse(http.StatusOK,"Admin authenticated successfully",admin,nil)
	c.JSON(http.StatusOK,successRes)
}