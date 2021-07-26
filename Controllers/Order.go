package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/bootcamp-microservice-api/Config"
	"github.com/bootcamp-microservice-api/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync"
	"time"
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

func ConcurrentOrder(c *gin.Context){
	go CreateOrder(c)
}

// CreateOrder ... Create student
func CreateOrder(c *gin.Context) {
	var order Models.Order
	c.BindJSON(&order)
	err := Models.CreateOrder(&order)
	if err!=nil{
		fmt.Println("error.")
	}
	var BuyQuantity = order.Quantity
	var product Models.Product
	var customer Models.User
	err_:=Models.GetProductByID(&product,strconv.FormatUint(uint64(order.Pid), 10))
	if err_ != nil {
		fmt.Println("Error error",http.StatusNotFound)
	}
	error_:=Models.GetUserByID(&customer,strconv.FormatUint(uint64(order.Cid), 10))
	if error_ != nil {
		fmt.Println("Error error",http.StatusNotFound)
	}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	if product.Quantity < BuyQuantity && customer.Active==false{
		order.Status = "Failed"
		p,_ := json.Marshal(product)
		fmt.Println(string(p),order.Pid)
	} else{
		order.Status = "Order Placed"
		NewQuantity := product.Quantity-BuyQuantity
		//Config.DB.Save(product)
		Config.DB.Model(&product).Where("pid = ?", product.Pid).Update("quantity", NewQuantity)
	}
	//if err != nil {
	//	fmt.Println(err.Error())
	//	c.AbortWithStatus(http.StatusNotFound)
	//} else {
	//	c.JSON(http.StatusOK,gin.H{
	//		"message":"order placed successfully",
	//	})
	//}
	mutex.Unlock()
	Config.DB.Model(&customer).Where("cid = ?", customer.Cid).Update("active", false)
	time.Sleep(100*time.Second)
	Config.DB.Model(&customer).Where("cid = ?", customer.Cid).Update("active", true)
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
