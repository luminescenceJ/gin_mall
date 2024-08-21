package dao

import (
	"fmt"
	"gin_mal_tmp/model"
)

func migration() {
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Favorite{},
			&model.Order{}, &model.Admin{}, &model.Address{},
			&model.Cart{}, &model.Category{}, &model.Carousel{},
			&model.Notice{}, &model.Notice{}, &model.Product{},
			&model.ProductImg{},
		)
	if err != nil {
		fmt.Println(err)
		return
	}
}
