package modle

import (
	"github.com/jinzhu/gorm"
)
//注意注意，这里不是指用户下单是指电影院的库存
type Moives struct {
	gorm.Model
	Moive_Name    string
	Moive_Time    string
	Moive_Positon string
	Moive_Num     int//库存
}



