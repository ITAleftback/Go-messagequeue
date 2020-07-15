package main

import (
	"summercourse3/modle"
	"summercourse3/mq"
)

func main()  {
	modle.InitDB()

	mq.OpenConsumer()
}
