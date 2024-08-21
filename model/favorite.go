package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"foreignkey:UserID"`
	UserID    uint    `gorm:"not null"`
	Product   Product `gorm:"foreignkey:ProductID"`
	ProductID uint    `gorm:"not null"`
	Seller    User    `gorm:"foreignkey:SellerID"`
	SellerID  uint    `gorm:"not null"`
}
