package main

import (
	"bubbon/dao"
	"bubbon/moudles"
	"bubbon/routers"
)

func main() {
	// 创建mysql数据库，连接数据库
	err := dao.InitSQL()
	if err != nil {
		panic(err)
	}
	defer dao.CloseDB()

	// 根据模型创建表结构
	dao.DB.AutoMigrate(&moudles.Todo{})

	// 创建路由
	route := routers.SetupRoute()

	route.Run(":8080")
}
