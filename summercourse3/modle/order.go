package modle

import (
	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	User_ID int
	Moive_ID int
}


func (orders *Order)Makeorders() error {
	return DB.Create(orders).Error
}

func Updatemoives(MoiveID int)error{
	return DB.Model(Moives{}).Where("id=?", MoiveID).Update("moive_num",gorm.Expr("moive_num-?",1)).Error
}

func (moives *Moives)Addmoives()error  {
	return DB.Create(moives).Error
}