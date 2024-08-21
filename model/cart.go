package model

import "gorm.io/gorm"

type Cart struct { //购物车
	gorm.Model
	UserID    uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	SellerID  uint `gorm:"not null"`
	Num       uint `gorm:"not null"`
	MaxNum    uint `gorm:"not null"`
	Check     uint // 是否支付
}
