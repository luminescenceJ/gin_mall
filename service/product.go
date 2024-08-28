package service

import (
	"context"
	"gin_mal_tmp/dao"
	"gin_mal_tmp/model"
	"gin_mal_tmp/pkg/e"
	"gin_mal_tmp/pkg/util"
	"gin_mal_tmp/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id" `
	Name          string `json:"name" form:"name" `
	CategoryId    uint   `json:"category_id" form:"category_id" `
	Title         string `json:"title" form:"title" `
	Info          string `json:"info" form:"info" `
	ImgPath       string `json:"img_path" form:"img_path" `
	Price         uint   `json:"price" form:"price" `
	DiscountPrice uint   `json:"discount" form:"discount" `
	OnSale        bool   `json:"on_sale" form:"on_sale" `
	Num           int    `json:"num" form:"num" `
	model.BasePage
}

func (s *ProductService) Create(c context.Context, uId uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(c)
	boss, _ = userDao.GetUserById(uId)

	//第一张图片作为封面
	tmp, _ := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uId, s.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          s.Name,
		CategoryId:    s.CategoryId,
		Title:         s.Title,
		Info:          s.Info,
		ImgPath:       path,
		Price:         s.Price,
		DiscountPrice: s.DiscountPrice,
		OnSale:        true,
		Num:           s.Num,
		SellerID:      uId,
		SellerName:    boss.UserName,
		SellerAvatar:  boss.Avatar,
	}

	productDao := dao.NewProductDao(c)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("Create Product api ", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocalStatic(tmp, uId, s.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.Error
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}
