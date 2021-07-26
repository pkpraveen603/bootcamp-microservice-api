package main
import (
	"fmt"
	"github.com/bootcamp-microservice-api/Config"
	"github.com/bootcamp-microservice-api/Models"
	"github.com/bootcamp-microservice-api/Routes"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{}, &Models.Product{}, &Models.Order{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}