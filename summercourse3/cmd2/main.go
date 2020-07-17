package main

import (
	"summercourse3/modle"
	"summercourse3/mq"
)

func main()  {
	modle.InitDB()
	//开启消费者进行消费
	mq.OpenConsumer()
}
