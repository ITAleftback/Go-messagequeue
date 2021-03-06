package mq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
	"summercourse3/message"
	"summercourse3/modle"
)

//拿出消息队列的数据将信息写入订单表，消费者
func OpenConsumer(){
	//连接队列
	conn,err:=amqp.Dial("amqp://guest:guest@localhost:5672/")
	message.HandleError(err,"Can't connect to MQ")
	defer conn.Close()
	amqpChannel,err:=conn.Channel()
	message.HandleError(err,"Can't create a amqpChannel")
	defer amqpChannel.Close()
	//想要连接的队列名，及其设置
	queue,err:=amqpChannel.QueueDeclare("orderlist",true,false,false,false,nil)
	message.HandleError(err,"Could not declare `add` queue")
	err = amqpChannel.Qos(1, 0, false)
	message.HandleError(err, "Could not configure QoS")

	messageChannel,err:=amqpChannel.Consume(queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	message.HandleError(err,"Could not register consumer")
	//用来关闭进程？
	stopChan :=make(chan bool)
	go func() {
		log.Printf("Consumer ready,PID:%d",os.Getpid())
		//遍历消息队列并拿出数据
		for d:=range messageChannel{
			log.Printf("Received a message:%s",string(d.Body))
			order:=&modle.Order{}
			err:=json.Unmarshal(d.Body,order)
			//此时，order就是消息队列拿出的数据
			if err != nil {
				log.Printf("Error decoding JSON:%s",err)
			}
			log.Printf("Order:%s",string(d.Body))
			//将从消息队列拿到的消息放进数据库
			orders:=modle.Order{
				User_ID:    order.User_ID,
				Moive_ID:   order.Moive_ID,
			}
			_ = orders.Makeorders()
			//拿出消息后，电影票的库存就会减一次
			_ = modle.Updatemoives(order.Moive_ID)


			if err:=d.Ack(false);err!=nil{
				log.Printf("Error acknowledgeing message : %s",err)
			}else{
				log.Printf("Acknowledged message")
			}
		}
	}()
	//关闭进程， 原理好像是channel阻塞？
	<-stopChan
}
