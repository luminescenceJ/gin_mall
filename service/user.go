package service

import (
	"context"
	"fmt"
	"gin_mal_tmp/conf"
	"gin_mal_tmp/dao"
	"gin_mal_tmp/model"
	"gin_mal_tmp/pkg/e"
	"gin_mal_tmp/pkg/util"
	"gin_mal_tmp/serializer"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"
)

type UserService struct {
	NickName string `json:"nick_name" form:"nick_name" binding:"required"`
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Key      string `json:"key" form:"key"` //前端验证 传输加密密文
}

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	// 1.绑定邮箱 2.解绑邮箱 3.修改密码
}
type ValidEmailService struct {
}
type ShowMoneyService struct {
	Key string `json:"key" form:"key"`
}

func (us *UserService) Register(ctx context.Context) serializer.Response {
	var user model.User
	var money string
	code := e.Success
	if us.Key == "" || len(us.Key) != 16 {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  "密钥长度不足",
		}
	}
	// 密码加密
	util.Encrypt.SetKey(us.Key)
	fmt.Println("加密的key", us.Key)
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
	money, err = util.Encrypt.AseEncrypt("10000")
	user = model.User{
		UserName: us.UserName,
		NickName: us.NickName,
		Status:   model.Activate,
		Avatar:   "avatar.jpg",
		Money:    money,
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

func (us *UserService) Update(ctx context.Context, uId uint) serializer.Response {
	var user *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if us.NickName != "" && us.UserName != "" {
		user.NickName = us.NickName
		user.UserName = us.UserName
	}
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

func (us *UserService) Post(ctx context.Context, uId uint, file multipart.File, fileSize int64) serializer.Response {
	var user *model.User
	var err error
	code := e.Success

	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	path, err := UploadAvatarToLocalStatic(file, uId, us.UserName)
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uId, user)
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}

func (ss *SendEmailService) Send(ctx context.Context, uId uint) serializer.Response {

	code := e.Success
	var address string
	var notice *model.Notice //模板通知（修改密码和邮箱绑定）
	token, err := util.GenerateEmailToken(uId, ss.OperationType, ss.Email, ss.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 模块化固定邮件格式
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetUserById(ss.OperationType)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = conf.ValidEmail + token
	mailStr := notice.Text
	mailText := strings.Replace(mailStr, "Email", address, -1)

	m := mail.NewMessage()
	m.SetHeader("From", conf.SmtpEmail)
	m.SetHeader("To", ss.Email)
	m.SetHeader("Subject", "admin")
	m.SetBody("text/html", mailText)

	d := mail.NewDialer(conf.SmtpHost, 465, conf.SmtpEmail, conf.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		//Data:   serializer.BuildUser(user),
	}
}

func (vs *ValidEmailService) Valid(ctx context.Context, token string) serializer.Response {
	var (
		userId        uint
		email         string
		password      string
		operationType uint
	)
	code := e.Success
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := util.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userId = claims.UserID
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.Success {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 验证成功 获取user信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.ErrorExistUserNotFound
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	if operationType == 1 {
		user.Email = email
	} else if operationType == 2 {
		user.Email = ""
	} else if operationType == 3 {
		err = user.SetPassword(password)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}

}

func (ss *ShowMoneyService) Show(ctx context.Context, uId uint) serializer.Response {
	var moneyInfo serializer.Money
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	moneyInfo, err = serializer.BuildMoney(user, ss.Key)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   moneyInfo,
	}
}
