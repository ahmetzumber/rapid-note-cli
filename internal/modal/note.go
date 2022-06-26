package modal

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	ID				int `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID 			int `gorm:"foreignKey:ID;references:UserID"`
	Data			string
}