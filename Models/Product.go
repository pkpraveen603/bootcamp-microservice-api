package Models
import (
	"fmt"
	"github.com/bootcamp-microservice-api/Config"
	_ "github.com/go-sql-driver/mysql"
)

// GetProducts Fetch all user data
func GetProducts(product *[]Product) (err error) {
	if err = Config.DB.Find(product).Error; err != nil {
		return err
	}
	return nil
}

// CreateProduct  ... Insert New data
func CreateProduct(product *Product) (err error) {
	if err = Config.DB.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// GetProductByID  ... Fetch only one user by Id
func GetProductByID(product *Product, id string) (err error) {
	if err = Config.DB.Where("pid = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

// UpdateProduct ... Update user
func UpdateProduct(product *Product, id string) (err error) {
	fmt.Println(product)
	Config.DB.Save(product)
	return nil
}

// DeleteProduct ... Delete user
func DeleteProduct(product *Product, id string) (err error) {
	Config.DB.Where("pid = ?", id).Delete(product)
	return nil
}