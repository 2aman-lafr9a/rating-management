package models

import (
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Rating struct {
	gorm.Model
	ID       string `gorm:"primaryKey;type:varchar(255);unique;autoIncrement:true"`
	OfferID  string `gorm:"type:varchar(255);not null"`
	PlayerID string `gorm:"type:varchar(255);not null"`
	Rating   int32  `gorm:"type:int;not null;default:0"`
}

func (r Rating) Error() string {
	err := r.Error()
	return err

}
