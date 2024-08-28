package main

import (
	"gin_mal_tmp/conf"
	"gin_mal_tmp/dao"
	"gin_mal_tmp/pkg/util"
	"gin_mal_tmp/routes"
)

func main() {
	conf.InitConfig()
	dao.InitMysql()
	util.InitLog()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
