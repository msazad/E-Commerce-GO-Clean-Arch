// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/msazad/go-Ecommerce/pkg/api"
	"github.com/msazad/go-Ecommerce/pkg/api/handlers"
	"github.com/msazad/go-Ecommerce/pkg/config"
	"github.com/msazad/go-Ecommerce/pkg/db"
	"github.com/msazad/go-Ecommerce/pkg/repository"
	"github.com/msazad/go-Ecommerce/pkg/usecase"
)

func InitializeAPI(cfg config.Config)(*api.ServerHTTP,error){
	gormDB,err:=db.ConnectDB(cfg)
	if err!=nil{
		return nil,err
	}

	offerRepository:=repository.NewOfferRepository(gormDB)
	offerUsecase:=usecase.NewOfferUsecase(offerRepository)
	offerHandler:=handlers.NewOfferHandler(offerUsecase)

	orderRepository:=repository.NewOrderRepository(gormDB)

	userRepository:=repository.NewUserRepository(gormDB)
	
	userUsecase:=usecase.NewUserUsecase(userRepository,offerRepository,orderRepository)

	couponRepository:=repository.NewCouponRepository(gormDB)

	orderUsecase:=usecase.NewOrderUsecase(orderRepository,userUsecase,couponRepository)
	orderHandler:=handlers.NewOrderHandler(orderUsecase)

	userHandler:=handlers.NewUserHandler(userUsecase)

	wishlistRepository:=repository.NewWishlistRepository(gormDB)
	wishlistUsecase:=usecase.NewWishlistUsecase(wishlistRepository)
	wishlistHandler:=handlers.NewWishlistHandler(wishlistUsecase)

	adminRepository :=repository.NewAdminRepository(gormDB)
	adminUsecase:=usecase.NewAdminUsecase(adminRepository)
	adminHandler:=handlers.NewAdminHandler(adminUsecase)

	

	

	inventoryRepository:=repository.NewInventoryRepository(gormDB)
	inventoryUsecase:=usecase.NewInventoryUsecase(inventoryRepository)
	inventoryHandler:=handlers.NewInventoryHandler(inventoryUsecase)

	categoryRepository:=repository.NewCategoryRepository(gormDB)
	categoryUsecase:=usecase.NewCategoryUsecase(categoryRepository)
	categoryHandler:=handlers.NewCategoryHandler(categoryUsecase)

	paymentRepo:=repository.NewPaymentRepository(gormDB)
	paymentUsecase:=usecase.NewPaymentUsecase(paymentRepo,userRepository)
	paymentHandler:=handlers.NewPaymentHandler(paymentUsecase)

	cartRepository:=repository.NewCartRepository(gormDB)
	cartUsecase:=usecase.NewCartUsecase(cartRepository,inventoryRepository,userUsecase,paymentUsecase)
	cartHandler:=handlers.NewCartHandler(cartUsecase)

	
	couponUsecase:=usecase.NewCouponUsecase(couponRepository)
	couponHandler:=handlers.NewCouponHandler(couponUsecase)


	otpRepository:=repository.NewOtpRepository(gormDB)
	otpUsecase:=usecase.NewOtpUsecase(cfg,otpRepository)
	otpHandler:=handlers.NewOtpHandler(otpUsecase)

	serverHTTP:=api.NewServerHttp(categoryHandler,inventoryHandler,userHandler,otpHandler,adminHandler,cartHandler,orderHandler,paymentHandler,wishlistHandler,offerHandler,couponHandler)

	return serverHTTP,nil
}