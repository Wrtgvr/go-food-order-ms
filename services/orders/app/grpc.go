package app

import (
	"log/slog"
	"net"

	"github.com/wrtgvr/go-food-order-ms/services/orders/handler"
	"github.com/wrtgvr/go-food-order-ms/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
	log  *slog.Logger
}

func newGRPCServer(addr string, log *slog.Logger) *gRPCServer {
	return &gRPCServer{
		addr: addr,
		log:  log,
	}
}

func (s *gRPCServer) Run() error {
	// listener and server
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	// services & handler
	ordersService := service.NewOrdersService()

	handlerDeps := &handler.HandlerDeps{
		OrdersService: ordersService,
		Log:           s.log,
	}

	handler.NewOrdersGrpcHandler(grpcServer, handlerDeps)

	// log
	s.log.Info("Launching server", slog.String("addr", s.addr))

	// serve
	return grpcServer.Serve(ln)
}
