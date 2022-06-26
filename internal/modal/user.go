package modal

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID				int `gorm:"primaryKey;autoIncrement:true;unique"`
	Username 		string
	Email	 		string
}