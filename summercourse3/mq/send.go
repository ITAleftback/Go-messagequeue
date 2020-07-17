package mq

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"summercourse3/message"
	"summercourse3/modle"
	"time"
)

//用来将传过来的参数放进消息队列
//用来将传过来的参数放进消息队列
func Makeorder(order modle.Order){
	//连接队列
	conn,err:=amqp.Dial("amqp://guest:guest@localhost:5672/")
	message.FailError(err,"Can't connect to MQ")
	defer conn.Close()
	amqpChannel,err:=conn.Channel()
	message.FailError(err,"Can't create a Channel")
	defer amqpChannel.Close()
	//想要连接的队列名 及其设置
	queue,err:=amqpChannel.QueueDeclare("orderlist",true,false,false,false,nil)
	message.FailError(err,"Could not declare queue")
	rand.Seed(time.Now().UnixNano())
	//序列化
	body,err:=json.Marshal(order)
	message.FailError(err,"Error encoding JSON")
	//将数据发送给队列，数据在body
	err=amqpChannel.Publish("",queue.Name,false,false,amqp.Publishing{

		ContentType:     "text/plain",
		DeliveryMode:    amqp.Persistent,
		Body:            body,
	})
	if err!=nil {
		log.Fatalf("Error publishing message:%s",err)
	}
	log.Printf("Makeorder:%s",string(body))

}