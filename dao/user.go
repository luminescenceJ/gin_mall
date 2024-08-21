package dao

import (
	"context"
	"gin_mal_tmp/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(name string) (user model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where("username=?", name).Find(&user).Error
	if err != nil {
		return
	}
	if user != (model.User{}) { //用户已存在
		return user, true, nil
	}
	return user, false, nil //可以创建
}

func (dao *UserDao) CreateUser(user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}
