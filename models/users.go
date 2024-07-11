package models

type User struct {
	ID          uint32  `gorm:"AUTO_INCREMENT" json:"id"`
	FirstName   string  `gorm:"type:varchar(100);not null" json:"first_name"`
	LastName    string  `gorm:"type:varchar(100);not null" json:"last_name"`
	NickName    string  `gorm:"type:varchar(100);not null" json:"nick_name"`
	Age         uint32  `json:"age"`
	Sex         string  `json:"sex"`
	Occupation  string  `json:"occupation"`
	Salary      float32 `json:"salary"`
	Email       string  `json:"email"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phone_number"`
}
