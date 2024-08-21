package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string
	Category      uint
	Title         string
	Info          string
	ImgPath       string
	Price         uint
	DiscountPrice uint
	OnSale        bool `gorm:"default:false"`
	Num           int
	SellerID      uint
	SellerName    string
	SellerAvatar  string
}
