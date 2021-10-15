package service

import (
	"context"
	rpc_package "grpc-example/api"
)

type HelloWordlService struct {


}

func (p *HelloWordlService) mustEmbedUnimplementedHelloWorldServiceServer() {
	panic("implement me")
}

func (p *HelloWordlService) SayHello(ctx context.Context, request *rpc_package.HelloRequest) (*rpc_package.HelloReply, error) {
	panic("implement me")
}

func (p *HelloWordlService) BidirectionalStream(server rpc_package.HelloWorldService_BidirectionalStreamServer) error {
	panic("implement me")
}

func NewHelloWorldService() *HelloWordlService {
	return &HelloWordlService{
	}
}

//func (p *HelloWordlService) SayHello(ctx context.Context, in *rpc_package.HelloRequest) (*rpc_package.HelloReply, error){
//	fmt.Println("receive req:",in)
//	return &rpc_package.HelloReply{Message:"success"},nil
//}
//
////双向流通信
//func (p *HelloWordlService) BidirectionalStream(stream rpc_package.HelloWorldService_BidirectionalStreamServer) error {
//	for{
//		n := 0
//		//server send
//		err := stream.Send(&rpc_package.SteamResponse{
//			OptCode:              int32(n),
//			RespCode:             0,
//			Output:               "server response",
//		})
//		if err != nil{
//			fmt.Println("server send error:",err)
//			return err
//		}
//		//server rec
//		r,err := stream.Recv()
//		if err == io.EOF{
//			fmt.Println("cliet disconnect,:",err)
//			return nil
//		}
//		if err != nil{
//			fmt.Println("client err:",err)
//			return  err
//		}
//		n++
//		fmt.Println("server rec info",r)
//	}
//}
