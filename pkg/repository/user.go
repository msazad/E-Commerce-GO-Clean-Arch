package repository

import (
	"errors"

	"github.com/msazad/go-Ecommerce/pkg/domain"
	"github.com/msazad/go-Ecommerce/pkg/repository/interfaces"
	"github.com/msazad/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: DB}
}

func (c *userDatabase) CheckUserAvailability(email string) bool {
	var count int
	query := `select count(*) from users where email=?`
	if err := c.DB.Raw(query, email).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0
}
func (c *userDatabase) UserBlockStatus(email string) (bool, error) {
	var permission bool
	err := c.DB.Raw("select permission from users where email=?", email).Scan(&permission).Error
	if err != nil {
		return false, err
	}
	return permission, nil
}

func (c *userDatabase) FindUserByEmail(user models.UserLogin) (models.UserResponse, error) {
	var user_details models.UserResponse

	err := c.DB.Raw(`SELECT * FROM users where email=? and permission=true
	`, user.Email).Scan(&user_details).Error

	if err != nil {
		return models.UserResponse{}, errors.New("error checking user details")
	}

	return user_details, nil
}

func (c *userDatabase) FindUserIDByOrderID(orderID int) (int, error) {

	var userID int

	err := c.DB.Raw(`
	SELECT user_id
	FROM orders where id=?
	`, orderID).Scan(&userID).Error

	if err != nil {
		return 0, errors.New("error checking user details")
	}
	return userID, nil
}

func (c *userDatabase) SignUp(user models.UserDetails) (models.UserResponse, error) {

	var userDetails models.UserResponse
	err := c.DB.Raw("INSERT INTO useres (name,email,password,phone,username)VALUES (?,?,?,?,?) RETURNING id,name,email,phone", user.Name, user.Email, user.Password, user.Phone, user.Username).Scan(&userDetails).Error

	if err != nil {
		return models.UserResponse{}, err
	}
	return userDetails, nil
}
func (i *userDatabase) AddAddress(id int, address models.AddAddress, result bool) error {
	err := i.DB.Exec(`
	INSERT INTO addresses (user_id,name,house_name,street,city,state,pin,"default")
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	RETURNING id`,
		id, address.Name, address, address.HouseName, address.Street, address.City, address.State, address.Pin, result).Error
	if err != nil {
		return errors.New("could not add address")
	}
	return nil
}

func (c *userDatabase) CheckIfFirstAddress(id int) bool {
	var count int
	// query:=fmt.Sprintf("select count(*) from addresses where user id=%s",id)
	if err := c.DB.Raw("select count(*) from addresses where user_id=?", id).Scan(&count).Error; err != nil {
		return false
	}
	// if count is greater than 0 that means the user already exist
	return count > 0
}
func (ad *userDatabase) GetAddresses(id int) ([]domain.Address, error) {

	var addresses []domain.Address

	if err := ad.DB.Raw("select * from addresses where user_id=?", id).Scan(&addresses).Error; err != nil {
		return []domain.Address{}, errors.New("error in getting addresses")
	}
	return addresses, nil
}

func (ad *userDatabase) GetUserDetails(id int) (models.UserResponse, error) {

	var details models.UserResponse

	if err := ad.DB.Raw("select id,name,username,email,phone from useres where id=?", id).Scan(&details).Error; err != nil {
		return models.UserResponse{}, errors.New("could not get user details")
	}

	return details, nil
}

func (i *userDatabase) ChangePassword(id int, password string) error {
	err := i.DB.Exec("UPDATE users SET password=? WHERE id=?", password, id).Error
	if err != nil {
		return errors.New("couldnt change password")
	}

	return nil
}

func (i *userDatabase) GetPassword(id int) (string, error) {

	var userPassword string
	err := i.DB.Raw("select password from users where id=?", id).Scan(&userPassword).Error
	if err != nil {
		return "", errors.New("couldnt get password")
	}
	return userPassword, nil
}

func (ad *userDatabase) FindIdFromPhone(phone string) (int, error) {

	var id int

	if err := ad.DB.Raw("select id from users where phone=?", phone).Scan(&id).Error; err != nil {
		return 0, err
	}

	return id, nil
}

func (i *userDatabase) EditName(id int, name string) error {
	err := i.DB.Exec(`update users set name=? where id=?`, name, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (i *userDatabase) EditEmail(id int, email string) error {
	err := i.DB.Exec(`update users set email=? where id=?`, email, id).Error
	if err != nil {
		return err
	}
	return nil
}

func(i *userDatabase)EditUsername(id int,username string)error{
	err:=i.DB.Exec(`update users set username=? where id=?`,username,id).Error
	if err!=nil{
		return err
	}
	return nil
}