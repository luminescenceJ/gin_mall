package main

import (
	"gin_mal_tmp/conf"
	"gin_mal_tmp/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
