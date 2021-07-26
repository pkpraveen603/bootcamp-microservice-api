package Models
type Product struct {
	Pid       			uint `json:"pid" gorm:"primaryKey;autoIncrement"`
	PricePerProduct     int  `json:"status"`
	Quantity  			int  `json:"quantity"`
}
func (p *Product) TableName() string {
	return "product"
}
