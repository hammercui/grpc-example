/*
@Desc : 2019/7/3 15:11 
@Version : 1.0.0
@Time : 2019/7/3 15:11 
@Author : hammercui
@File : main
@Company: Sdbean
*/
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-example/api"
	"log"
	"time"
)
const(
	address = "localhost:9010"
)
func main()  {
	conn,err := grpc.Dial(address,grpc.WithInsecure())
	if err !=nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := rpc_package.NewHelloWorldServiceClient(conn)

	for i:=0; i<5; i++ {
		go SendTest(client)
	}

	time.Sleep(5*time.Second)
}

func SendTest(client rpc_package.HelloWorldServiceClient){
	start := time.Now()
	//send request
	resp, err := client.SayHello(context.Background(), &rpc_package.HelloRequest{Name:"i am client"})
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	fmt.Println("receive",resp,time.Since(start))
}