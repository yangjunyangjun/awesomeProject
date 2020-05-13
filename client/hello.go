package main

import (
	"awesomeProject/proto"
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
)
func main(){
	//创建一个新的服务
	server := micro.NewService(micro.Name("getHello"))
	//初始化
	server.Init()
	hello := TestHello.NewHelloService("hello",server.Client())
	response,err:=hello.Hello(context.TODO(),&TestHello.Hrequest{
		Name:                 "yangjun",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(response.Msg)
}