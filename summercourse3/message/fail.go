package message

import "log"
//失误的消息
func FailError(err error,msg string){
	if err!=nil {
		log.Fatalf("%s:%s",msg,err)
	}
}

func HandleError(err error,msg string){
	if err!=nil {
		log.Fatalf("%s:%s",msg,err)
	}
}
