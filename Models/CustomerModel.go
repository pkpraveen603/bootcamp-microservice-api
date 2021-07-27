package Models
type User struct {
	Cid      uint64   `json:"cid" gorm:"primaryKey;autoIncrement:true"`
	Name     string `json:"name"`
	Active   bool    `json:"active"`
}
func (b *User) TableName() string {
	return "user"
}