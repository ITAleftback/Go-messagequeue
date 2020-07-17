package modle

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	User_ID int
	Moive_ID int
}

//下单数据库数据库操作
func (orders *Order)Makeorders() error {
	return DB.Create(orders).Error
}
//更新数据库电影的操作
func Updatemoives(MoiveID int)error{
	return DB.Model(Moives{}).Where("id=?", MoiveID).Update("moive_num",gorm.Expr("moive_num-?",1)).Error
}
//添加电影的数据库操作
func (moives *Moives)Addmoives()error  {
	return DB.Create(moives).Error
}