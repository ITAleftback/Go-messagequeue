package message

import "log"

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
