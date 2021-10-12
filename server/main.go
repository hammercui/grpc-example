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
	"context"
	"fmt"
	"grpc-example/api"
	"log"
	"net"
	"google.golang.org/grpc"
)

const(
	address = "localhost:9010"
)
type service struct {

}

func (p *service) SayHello(ctx context.Context, in *rpc_package.HelloRequest) (*rpc_package.HelloReply, error){
	fmt.Println("receive req:",in)
	return &rpc_package.HelloReply{Message:"success"},nil
}

func main(){
	lis,err := net.Listen("tcp",address)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("list server success!",address)
	s := grpc.NewServer()
	//注册service实现
	rpc_package.RegisterHelloWorldServiceServer(s,&service{})
	fmt.Println("register service success!")
	defer lis.Close()
	err = s.Serve(lis)
	if err !=nil{
		log.Fatalf("failed to run server : %v", err)
	}else{
		fmt.Println("run server success!",address)
	}
}
