package model

type BasePage struct { //数据库分页
	PageNum  int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}
