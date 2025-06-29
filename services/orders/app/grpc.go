package app

import (
	"log"
	"net"

	"github.com/wrtgvr/go-food-order-ms/services/orders/handler"
	"github.com/wrtgvr/go-food-order-ms/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func newGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

func (s *gRPCServer) Run() error {
	// listener and server
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	// services
	ordersService := service.NewOrdersService()

	handler.NewOrdersGrpcHandler(grpcServer, ordersService)

	// log
	log.Printf("Server starting on %s", s.addr)

	// serve
	return grpcServer.Serve(ln)
}
