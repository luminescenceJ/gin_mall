package serializer

import (
	"gin_mal_tmp/model"
	"gin_mal_tmp/pkg/util"
)

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"username" form:"username"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(item *model.User, key string) (Money, error) {
	util.Encrypt.SetKey(key)
	money, err := util.Encrypt.AseDecrypt(item.Money)
	if err != nil {
		return Money{}, err
	}
	return Money{
		UserId:    item.ID,
		UserName:  item.UserName,
		UserMoney: money,
	}, nil
}
