package Controllers
import (
	"fmt"
	"github.com/bootcamp-microservice-api/Models"
	"github.com/gin-gonic/gin"
	"net/http"
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
