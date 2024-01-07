package repository

import (
	"time"

	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type orderRepository struct{
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB)interfaces.OrderRepository{
	return &orderRepository{
		DB: db,
	}
}

func (o *orderRepository)GetOrders(id,page,limit int)([]domain.Order,error){

	if page ==0{
		page=1
	}
	if limit==0{
		limit=10
	}
	offset:=(page-1)*limit
	var orders []domain.Order

	if err:=o.DB.Raw("select * from orders where user_id=? limit ? offset ?",id,limit,offset).Scan(&orders).Error;err!=nil{
		return []domain.Order{},err
	}
	return orders,nil
}

func (o *orderRepository)GetOrdersInRange(startDate,endDate time.Time)([]domain.Order,error){
	var orders []domain.Order

	// Execute the query to get orders within the specified time range
	if err:=o.DB.Raw("SELECT * FROM orders WHERE ordered_at BETWEEN ? AND ?",startDate,endDate).Scan(&orders).Erro;err!=nil{
		return []domain.Order{},err
	}
	return orders,nil
}

func (o *orderRepository)GetProductsQuantity()([]domain.ProductReport,error){

	var products []domain.ProductReport

	if err:=o.DB.Raw("select inventory_id,quantity from order_items").Scan(&products).Error;err!=nil{
		return []domain.ProductReport{},err
	}
	return products,nil
}

