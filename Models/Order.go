package Models
import (
	"fmt"
	"github.com/bootcamp-microservice-api/Config"
	_ "github.com/go-sql-driver/mysql"
)

// GetAllOrders Fetch all user data
func GetAllOrders(order *[]Order) (err error) {
	if err = Config.DB.Find(order).Error; err != nil {
		return err
	}
	return nil
}

// CreateOrder ... Insert New data
func CreateOrder(order *Order) (err error) {
	if err = Config.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

// GetOrderForCustomerID ... Fetch only one user by Id
func GetOrderForCustomerID(order *Order, id string) (err error) {
	if err = Config.DB.Where("cid = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// UpdateOrder ... Update user
func UpdateOrder(order *Order, id string) (err error) {
	fmt.Println(order)
	Config.DB.Save(order)
	return nil
}

// GetOrderByID  ... Fetch only one user by Id
func GetOrderByID(order *Order, id string) (err error) {
	if err = Config.DB.Where("oid = ?", id).First(order).Error; err != nil {
		return err
	}
	return nil
}
////DeleteUser ... Delete user
//func DeleteUser(user *User, id string) (err error) {
//	Config.DB.Where("id = ?", id).Delete(user)
//	return nil
//}
