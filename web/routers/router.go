package routers

import (
	"bubbon/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	//添加静态文件映射
	r.Static("static", "F:\\web3\\training\\bubbon\\static")
	r.Static("templates", "F:\\web3\\training\\bubbon\\templates")

	// 加载和解析模板
	r.LoadHTMLFiles("F:\\web3\\training\\bubbon\\templates\\*")

	// root界面
	r.GET("/", controller.IndexHandle)

	// 设置分组
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.CreateATodo)
		v1Group.GET("/todo", controller.GetTodoList)
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}

	return r
}
