package Models
type Product struct {
	Pid       			uint `json:"pid" gorm:"primaryKey;autoIncrement"`
	ProductName         string `json:"name"`
	PricePerProduct     int  `json:"price"`
	Quantity  			int  `json:"quantity"`
}
func (p *Product) TableName() string {
	return "product"
}
