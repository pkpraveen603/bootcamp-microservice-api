package Models
type Product struct {
	Pid       			uint `json:"pid" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	ProductName         string `json:"name"`
	PricePerProduct     int  `json:"price"`
	Quantity  			int  `json:"quantity"`
}
func (p *Product) TableName() string {
	return "product"
}
