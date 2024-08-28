package model

import (
	"gin_mal_tmp/cache"
	"gin_mal_tmp/conf"
	"gorm.io/gorm"
	"strconv"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
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

func (product *Product) View() uint64 {
	conf.Redis()
	countStr, _ := conf.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (product *Product) AddView() {
	conf.RedisClient.Incr(cache.ProductViewKey(product.ID))
	conf.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
