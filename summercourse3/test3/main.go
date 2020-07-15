package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

//这个主函数是我拿来测试高并发的,cmd和cmd2的主函数

func main(){
	for i:=0;i<50 ;i++  {
		Gopost()
		time.Sleep(10 *time.Millisecond)
	}
}

func Gopost(){
	go func() {
		resp, err := http.PostForm("http://localhost:8080/makeorder", url.Values{"user_ID": {"10086"}, "moive_ID": {"3"}})

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
	}()


}


