package main

import (
	"template/conf"
	"template/server"
)

// @title 测试模板
// @version 1.0
// @description 这只是一个测试模板
// @contact.name JunSi Chen
// @contact.name ZhengHao Liu
// @host 9090
// @BasePath /api
func main() {
	if err := conf.Init(); err != nil {
		panic(err)
	}

	s := server.NewServer()
	if err := s.Init(conf.Conf); err != nil {
		panic(err)
	}
	err := s.Run()
	if err != nil {
		panic(err)
	}
}
