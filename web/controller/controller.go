package controller

import (
	"bubbon/moudles"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandle(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	// 先获取前端创建请求的参数信息
	var todo moudles.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
		return
	}
	// 存入数据库
	if err := moudles.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": todo})
	}

}

func GetTodoList(c *gin.Context) {
	allTodo, err := moudles.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})

	} else {
		c.JSON(http.StatusOK, gin.H{"msg": allTodo})
	}

}

func UpdateATodo(c *gin.Context) {
	// 获取指定更新的id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"err": "无效的id!"})
		return
	}
	// 查找该id数据
	todo, err := moudles.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err = c.BindJSON(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	// 将新值更新入库
	if err = moudles.UpdateATodo(todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": todo})
	}
}

func DeleteATodo(c *gin.Context) {
	//获取id
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"err": "无效id"})
		return
	}
	if err := moudles.DeleteATodo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("id: %s,删除成功", id)})
	}
}
