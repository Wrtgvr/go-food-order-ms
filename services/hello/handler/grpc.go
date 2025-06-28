package handler

import (
	"context"

	wrt_hello_v1 "github.com/wrtgvr/go-food-order-ms/services/common/genproto/hello"
	"github.com/wrtgvr/go-food-order-ms/services/hello/types"
	"google.golang.org/grpc"
)

type HelloGrpcHandler struct {
	helloService types.HelloService
	wrt_hello_v1.UnimplementedHelloServiceServer
}

func NewHelloGrpcHandler(grpc *grpc.Server, helloService types.HelloService) {
	h := &HelloGrpcHandler{
		helloService: helloService,
	}

	wrt_hello_v1.RegisterHelloServiceServer(grpc, h)
}

func (h *HelloGrpcHandler) SayHello(ctx context.Context, req *wrt_hello_v1.SayRequest) (*wrt_hello_v1.SayResponse, error) {
	name := req.GetName()

	res := h.helloService.SayHello(name)

	return &wrt_hello_v1.SayResponse{
		Message: res,
	}, nil
}
