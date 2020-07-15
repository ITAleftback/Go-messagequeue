package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"summercourse3/modle"
	"summercourse3/mq"
)
//加电影
func Addmovies(c *gin.Context){
	moive_Name:=c.PostForm("moive_Name")
	moive_Time:=c.PostForm("moive_Time")
	moive_Position:=c.PostForm("moive_Position")
	//添加多少库存
	moive_Num:=c.PostForm("moive_Num")
	moiveNum,_:=strconv.Atoi(moive_Num)

	moive:=modle.Moives{

		Moive_Name:    moive_Name,
		Moive_Time:    moive_Time,
		Moive_Positon: moive_Position,
		Moive_Num:     moiveNum,
	}
	moive.Addmoives()
	c.JSON(200,gin.H{"status":200,"info":"success"})

}
//下单的接口
func Makeorder(c *gin.Context){
	//输入你想订票的ID信息等等
	user_ID:=c.PostForm("user_ID")

	moive_ID:=c.PostForm("moive_ID")
	userID,_:=strconv.Atoi(user_ID)
	moiveID,_:=strconv.Atoi(moive_ID)
	order:=modle.Order{
		User_ID: userID ,
		Moive_ID: moiveID,
	}
	mq.Makeorder(order)


	c.JSON(200,gin.H{"status":200,"info":"success"})
}