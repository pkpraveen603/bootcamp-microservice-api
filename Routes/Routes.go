package Routes
import (
	"github.com/bootcamp-microservice-api/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//channel1 := make(chan bool)
	grp1 := r.Group("/customer-api")
	{
		grp1.GET("user", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
	}
	grp2 := r.Group("/retailer-api")
	{
		grp2.GET("product", Controllers.GetProducts)
		grp2.POST("product", Controllers.CreateProduct)
		grp2.GET("product/:id", Controllers.GetProductByID)
		grp2.PATCH("product/:id", Controllers.UpdateProduct)
		grp2.DELETE("product/:id", Controllers.DeleteProduct)
		grp2.PATCH("order/:id", Controllers.UpdateOrder)
	}
	grp3 := r.Group("/order-api")
	{
		grp3.GET("order", Controllers.GetAllOrders)
		grp3.POST("order", Controllers.ConcurrentOrder)
		grp3.GET("order/:id", Controllers.GetOrderForCustomerID)
		//grp2.DELETE("order/:id", Controllers.DeleteProduct)
	}

	return r
}