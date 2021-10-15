/*
@Desc : 2019/7/3 15:16 
@Version : 1.0.0
@Time : 2019/7/3 15:16 
@Author : hammercui
@File : main.go
@Company: Sdbean
*/
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-example/api"
	"grpc-example/server/service"
	"log"
	"net"
)

const(
	address = "localhost:9010"
)



func main(){
	lis,err := net.Listen("tcp",address)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("list server success!",address)
	s := grpc.NewServer()

	//注册service实现
	rpc_package.RegisterHelloWorldServiceServer(s, service.NewHelloWorldService())


	fmt.Println("register service success!")
	defer lis.Close()
	err = s.Serve(lis)
	if err !=nil{
		log.Fatalf("failed to run server : %v", err)
	}else{
		fmt.Println("run server success!",address)
	}
}
