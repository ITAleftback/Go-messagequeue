package main

import (
	"github.com/gin-gonic/gin"
	"summercourse3/controller"
	"summercourse3/modle"
)
func main(){
	modle.InitDB()

	router:=gin.Default()
	//这个函数是我拿来加电影的，方便后面的程序找到相应的电影并把库存-1
	router.POST("addmoives",controller.Addmovies)

	router.POST("/makeorder",controller.Makeorder)
	router.Run(":8080")


}