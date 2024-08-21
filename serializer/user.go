package serializer

import (
	"gin_mal_tmp/conf"
	"gin_mal_tmp/model"
)

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Type      int    `json:"type"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:        user.ID,
		Username:  user.UserName,
		Nickname:  user.NickName,
		Email:     user.Email,
		Status:    user.Status,
		Avatar:    conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
