package service

import (
	"context"
	"fmt"
	rpc_package "grpc-example/api"
	"io"
	"log"
	"sync"
)

type HelloWorldService struct {
	//对客户端广播map
	broadcastMap     map[uint64]rpc_package.HelloWorldService_BidirectionalStreamServer
	broadcastMapLock sync.Mutex
}

func NewHelloWorldService() *HelloWorldService {
	return &HelloWorldService{
		broadcastMap:     make(map[uint64]rpc_package.HelloWorldService_BidirectionalStreamServer),
		broadcastMapLock: sync.Mutex{},
	}
}

func (p *HelloWorldService) mustEmbedUnimplementedHelloWorldServiceServer() {
	panic("implement me")
}

func (p *HelloWorldService) SayHello(ctx context.Context, request *rpc_package.HelloRequest) (*rpc_package.HelloReply, error) {
	fmt.Println("receive req:", request)
	return &rpc_package.HelloReply{Message: "success"}, nil
}

func (p *HelloWorldService) BidirectionalStream(stream rpc_package.HelloWorldService_BidirectionalStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("收到客户端通过context发出的终止信号")
			return ctx.Err()
		default:
			//接收客户端消息
			r, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("cliet disconnect,:", err)
				uidVal := p.getSessionValue(ctx,"uid")
				uid := uidVal.(uint64)
				p.handleClientLeave(uid)
				return nil
			}
			if err != nil {
				log.Println("接收数据出错:", err)
				uidVal := p.getSessionValue(ctx,"uid")
				uid := uidVal.(uint64)
				p.handleClientLeave(uid)
				return err
			}
			//同步发送客户端消息
			p.handleRecRequest(r,stream)
		}
	}
}

func (p *HelloWorldService) handleRecRequest(request *rpc_package.StreamRequest,stream rpc_package.HelloWorldService_BidirectionalStreamServer) {
	if stream == nil{
		return
	}
	//	if err != nil{
	//		fmt.Println("client err:",err)
	//		return err
	//	}
	//	fmt.Println("server rec data :",r)
	//	//server受到消息胡发送send
	//	err = stream.Send(&rpc_package.SteamResponse{
	//		OptCode:              int32(r.OptCode),
	//		RespCode:             0,
	//		Output:               fmt.Sprintf("server response,cient data:%+v",r),
	//	})
	//	if err != nil{
	//		fmt.Println("server send error:",err)
	//	}

}

func (p *HelloWorldService) handleClientLeave(uid uint64) {

}

func (p *HelloWorldService) setSessionValue()  {

}
func (p *HelloWorldService) getSessionValue(ctx context.Context,key string) interface{}  {
	return 1
}

func (p *HelloWorldService) addBroadcastListener(uid uint64, stream rpc_package.HelloWorldService_BidirectionalStreamServer) {
	p.broadcastMapLock.Lock()
	defer p.broadcastMapLock.Unlock()
	p.broadcastMap[uid] = stream
}

func (p *HelloWorldService) removeBroadcastListener(uid uint64, stream rpc_package.HelloWorldService_BidirectionalStreamServer) {
	p.broadcastMapLock.Lock()
	defer p.broadcastMapLock.Unlock()
	delete(p.broadcastMap,uid)
}
