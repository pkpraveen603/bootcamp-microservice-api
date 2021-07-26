package Controllers
import (
	"encoding/json"
	"fmt"
	"github.com/bootcamp-microservice-api/Config"
	"github.com/bootcamp-microservice-api/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetAllOrders ... Get all students
func GetAllOrders(c *gin.Context) {
	var order []Models.Order
	err := Models.GetAllOrders(&order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

// CreateOrder ... Create student
func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.CreateOrder(&order)
	var BuyQuantity = order.Quantity
	var product Models.Product
	err_:=Models.GetProductByID(&product,strconv.FormatUint(uint64(order.Pid), 10))
	if err_ != nil {
		fmt.Println("Error error",http.StatusNotFound)
	}
	if product.Quantity < BuyQuantity{
		order.Status = "Failed"
		p,_ := json.Marshal(product)
		fmt.Println(string(p),order.Pid)
	} else{
		order.Status = "Order Placed"
		product.Quantity = product.Quantity-BuyQuantity
		Config.DB.Save(product)
	}

	//Models.Get
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

// GetOrderForCustomerID ... Get student by id
func GetOrderForCustomerID(c *gin.Context) {
	id := c.Params.ByName("cid")
	var order Models.Order
	err := Models.GetOrderForCustomerID(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

// UpdateOrder ... Update the student information
func UpdateOrder(c *gin.Context) {
	var order Models.Order
	id := c.Params.ByName("id")
	err := Models.GetOrderByID(&order, id)
	if err != nil {
		c.JSON(http.StatusNotFound, order)
	}
	error_ := c.BindJSON(&order)
	if error_ != nil {
		fmt.Println(error_.Error())
	}
	err = Models.UpdateOrder(&order, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

// DeleteProduct ... Delete the student record
//func DeleteProduct(c *gin.Context) {
//	var product Models.Product
//	id := c.Params.ByName("id")
//	err := Models.DeleteUser(&product, id)
//	if err != nil {
//		c.AbortWithStatus(http.StatusNotFound)
//	} else {
//		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
//	}
//}
