package Models
type User struct {
	Cid      uint   `json:"cid" gorm:"primaryKey;autoIncrement""`
	Name     string `json:"name"`
}
func (b *User) TableName() string {
	return "user"
}