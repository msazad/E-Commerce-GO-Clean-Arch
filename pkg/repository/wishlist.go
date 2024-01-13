package repository

import (
	"github.com/msazad/go-Ecommerce/pkg/repository/interfaces"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type WishlistRepository struct{
	DB *gorm.DB
}

func NewWishlistRepository(DB *gorm.DB)interfaces.WishlistRepository{
	return &WishlistRepository{DB}
}

func (w *WishlistRepository)GetWishlistID(id int)(int,error){

	var wishlistID int

	if err:=w.DB.Raw("select id from wishlist where user_id=?",id).Scan(&wishlistID).Error;err!=nil{
		return 0,err
	}

	return wishlistID,nil
}

func (w *WishlistRepository)GetWishlist(id int)([]models.GetWishlist,error){

	var wishlist []models.GetWishlist

	if err:=w.DB.Raw("select wishlist.user_id,categories.category,inventories.product_name,inventories.price from wishlists join wishlist_items on iwshlist_items.inventory_id=inventories.id join categories on inventories.category_id=categories.id where wishlists.user_id=?",id).Scan(&wishlist).Error;err!=nil{
		return []models.GetWishlist{},err
	}
	return wishlist,nil
}

func (w *WishlistRepository)GetProductsInWishlist(wishlistID int)([]int,error){

	var WishlistProducts []int
	if err:=w.DB.Raw("select inventory_id from wishlist_items where wishlist_id=?",wishlistID).Scan(&WishlistProducts).Error;err!=nil{
		return []int{},err
	}
	return WishlistProducts,nil
}

func (w *WishlistRepository)FindProductNames(inventory_id int)(string,error){

	var product_name string

	if err:=w.DB.Raw("select product_name from inventories where id=?",inventory_id).Scan(&product_name).Error;err!=nil{
		return "",err
	}

	return product_name,nil
}

func (w *WishlistRepository)FindPrice(inventory_id int)(float64,error){

	var price float64

	if err:=w.DB.Raw("select price from inventories where id=?",inventory_id).Scan(&price).Error;err!=nil{
		return 0,err
	}
	return price,nil
}

func (w *WishlistRepository)FindCategory(inventory_id int)(string,error){

	var category string

	if err:=w.DB.Raw("select categories.category from inventories join categories on inventories.category_id=categories.id where inventories.id=?",inventory_id).Scan(&category).Error;err!=nil{
		return "",err
	}

	return category,nil
}
//updation
func (w *WishlistRepository)RemoveFromWishlist(wishlistID int,inventoryID int)error{
	if err:=w.DB.Exec(`delete from wishlist_items where wishlist_id=? and inventory_id=?`,wishlistID,inventoryID).Error;err!=nil{
		return err
	}
	return nil
}