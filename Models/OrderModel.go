package Models
type Order struct {
	Oid       uint   `json:"oid" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	Cid       uint   `json:"cid"`
	Pid       uint   `json:"pid"`
	Status    string `json:"status"`
	Quantity  int    `json:"quantity"`
}
func (o *Order) TableName() string {
	return "orders"
}