package service

import (
	"context"
	"gin_mal_tmp/dao"
	"gin_mal_tmp/pkg/e"
	"gin_mal_tmp/pkg/util"
	"gin_mal_tmp/serializer"
)

type CarouselService struct {
}

func (s *CarouselService) List(ctx context.Context) serializer.Response {
	carouselDao := dao.NewCarouselDao(ctx)
	code := e.Success
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		util.LogrusObj.Infoln("err", err)
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
