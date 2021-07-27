package Controllers

import (
	"fmt"
	"github.com/bootcamp-microservice-api/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func ConcurrentUpdate(c *gin.Context){
	cc := c.Copy()
	//GetError := false
	//userReturned := Models.User{}
	ChannelErrorUp := make(chan bool,100)
	ChannelProductUp := make(chan Models.Product,100)
	go UpdateProduct(cc, ChannelErrorUp, ChannelProductUp)
	ans := <-ChannelErrorUp
	p := <-ChannelProductUp
	if ans {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, p)
	}
}

func ConcurrentDelete(c *gin.Context){
	cc := c.Copy()
	//GetError := false
	//userReturned := Models.User{}
	ChannelErrorDel := make(chan bool,100)
	//ChannelProductDel := make(chan Models.Product,100)
	go DeleteProduct(cc, ChannelErrorDel)
	ans := <-ChannelErrorDel
	//p := <-ChannelProductDel
	if ans {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":"Product deleted successfully",
		})
	}
}

// GetProducts ... Get all products
func GetProducts(c *gin.Context) {
	var product []Models.Product
	err := Models.GetProducts(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// CreateProduct ... Create new product
func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":product.Pid,
			"product_name":product.ProductName,
			"price":product.PricePerProduct,
			"quantity":product.Quantity,
			"message":"Product added successfully",
		})
	}
}

// GetProductByID ... Get p by id
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

// UpdateProduct ... Update the prod information
func UpdateProduct(c *gin.Context, ChannelErrorUp chan bool, ChannelProductUp chan Models.Product) {
	var product Models.Product
	mutex := &sync.Mutex{}
	mutex.Lock()
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		ChannelErrorUp<-true
		ChannelProductUp<-product
		//c.JSON(http.StatusNotFound, product)
	}
	error_ := c.BindJSON(&product)
	if error_ != nil {
		fmt.Println(error_.Error())
	}
	err = Models.UpdateProduct(&product, id)
	if err != nil {
		//c.AbortWithStatus(http.StatusNotFound)
		ChannelErrorUp<-true
		ChannelProductUp<-product
	} else {
		//c.JSON(http.StatusOK, product)
		ChannelErrorUp<-false
		ChannelProductUp<-product
	}
	mutex.Unlock()
}

// DeleteProduct ... Delete the record
func DeleteProduct(c *gin.Context, ChannelErrorDel chan bool) {
	var product Models.Product
	mutex := &sync.Mutex{}
	mutex.Lock()
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&product, id)
	if err != nil {
		//c.AbortWithStatus(http.StatusNotFound)
		ChannelErrorDel<-true
	} else {
		ChannelErrorDel<-false
		//c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
	mutex.Unlock()
}
