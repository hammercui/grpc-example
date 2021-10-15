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
	"io"
	"log"
	"time"
)

const (
	address = "localhost:9010"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := rpc_package.NewHelloWorldServiceClient(conn)

	//for i:=0; i<5; i++ {
	//	go SendTest(client)
	//}
	HandleStream(client)
	//time.Sleep(5*time.Second)
}

func SendTest(client rpc_package.HelloWorldServiceClient) {
	start := time.Now()
	//send request
	resp, err := client.SayHello(context.Background(), &rpc_package.HelloRequest{Name: "i am client"})
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	fmt.Println("receive", resp, time.Since(start))
}

func HandleStream(client rpc_package.HelloWorldServiceClient) {
	stream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		fmt.Println("err", err)
		return
	}
	for n := 0; n <= 6; n++ {
		err = stream.Send(&rpc_package.StreamRequest{
			OptCode: int32(n),
			Input:   "client input data",
		})
		if err != nil {
			fmt.Println("err", err)
			return
		}
	}

	defer stream.CloseSend()
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("io.EOF")
		}
		if err != nil {
			log.Println("err", err)
			return
		}
		log.Println("rec response code,", resp.OptCode)
		log.Println("rec response,", resp)
	}

}
