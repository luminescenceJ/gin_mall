package serializer

import (
	"gin_mal_tmp/conf"
	"gin_mal_tmp/model"
)

type Product struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"image"`
	Price         uint   `json:"price"`
	DiscountPrice uint   `json:"discount_price"`
	View          uint64 `json:"view"` //浏览次数
	CreateAt      int64  `json:"create_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(item *model.Product) Product {
	return Product{
		Id:            item.ID,
		Name:          item.Name,
		CategoryId:    item.CategoryId,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          item.View(),
		CreateAt:      item.CreatedAt.Unix(),
		Num:           item.Num,
		OnSale:        item.OnSale,
		BossId:        item.SellerID,
		BossName:      item.SellerName,
		BossAvatar:    item.SellerAvatar,
	}
}
