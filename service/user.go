package service

import (
	"context"
	"gin_mal_tmp/dao"
	"gin_mal_tmp/model"
	"gin_mal_tmp/pkg/e"
	"gin_mal_tmp/pkg/util"
	"gin_mal_tmp/serializer"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name" binding:"required"`
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Key      string `json:"key" form:"key"` //前端验证 传输加密密文

}

func (us *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	code := e.Success
	if us.Key == "" || len(us.Key) != 6 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}
	// 密码加密
	util.Encrypt.SetKey(us.Key)
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(us.UserName)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist { // 已经存在该名字
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user = model.User{
		Username: us.UserName,
		NickName: us.NickName,
		Status:   model.Activate,
		Avatar:   "avatar.jpg",
		Money:    util.Encrypt.AseEncoding("10000"),
	}

	//password encryption
	if err = user.SetPassword(us.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	//create user
	err = userDao.CreateUser(&user)
	if err != nil {
		code = e.Error
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (us *UserService) Login(ctx context.Context) serializer.Response {
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, exist, err := userDao.ExistOrNotByUserName(us.UserName)
	if !exist || err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "user not exist ,plz sign up",
		}
	}
	if !user.CheckPassword(us.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "wrong password ,try login in again",
		}
	}
	// token authorization
	token, err := util.GenerateToken(user.ID, us.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "wrong password ,try login in again",
		}
	}

	return serializer.Response{
		Status: code,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(&user),
			Token: token,
		},
		Msg: e.GetMsg(code),
	}
}
