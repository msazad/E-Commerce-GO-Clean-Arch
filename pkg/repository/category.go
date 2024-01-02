package repository

import (
	"errors"

	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)


type categoryRepository struct{
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB)interfaces.CategoryRepository{
	return &categoryRepository{DB}
}
func (p *categoryRepository)AddCategory(cat string)(domain.Category,error){
	var b string
	err:=p.DB.Raw("INSERT INTO categories (category) VALUES (?) RETURNING category",cat).Scan(&b).Error
	if err!=nil{
		return domain.Category{},err
	}
	var categoryResponse domain.Category
	err=p.DB.Raw(`
	SELECT
	p.id,
	p.category
	FROM
	    categories p
	WHERE
	     p.caregory=?
		 `,b).Scan(&categoryResponse).Error
	if err!=nil{
		return domain.Category{},err
	}
	return categoryResponse,nil
}

func (p *categoryRepository)CheckCategory(current string)(bool,error){
	var i int
	err:=p.DB.Raw("SELECT COUNT(*) FROM categories WHERE category=?",current).Scan(&i).Error
	if err !=nil{
		return false,err
	}
	if i==0{
		return false,err
	}
	return true,err
}

func (p *categoryRepository)UpdateCategory(current,new string)(domain.Category,error){
	//Check the database connection
	if p.DB==nil{
		return domain.Category{},errors.New("database conncetion is nil")
	}

	//Update the category
	if err:=p.DB.Exec("UPDATE categories SET category=? WHERE category =?",new,current).Error;err!=nil{
		return domain.Category{},err
	}
	
}