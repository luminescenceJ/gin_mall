package model

import "gorm.io/gorm"

type Carousel struct { // 轮播图
	gorm.Model
	ImgPath   string
	ProductId uint `gorm:not null`
}
