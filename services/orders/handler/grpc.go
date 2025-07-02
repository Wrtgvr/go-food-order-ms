package handler

import (
	"context"

	wrt_orders_v1 "github.com/wrtgvr/go-food-order-ms/services/common/genproto/orders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrdersGrpcHandler struct {
	ordersService OrdersService
	wrt_orders_v1.UnimplementedOrdersServiceServer
}

func NewOrdersGrpcHandler(grpcServer *grpc.Server, ordersService OrdersService) {
	h := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	wrt_orders_v1.RegisterOrdersServiceServer(grpcServer, h)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *wrt_orders_v1.CreateOrderRequest) (*wrt_orders_v1.CreateOrderResponse, error) {
	createOrderParams := toDomainCreateOrder(req)

	if err := createOrderParams.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "request validation failed: %v", err)
	}

	err := h.ordersService.CreateOrder(ctx, createOrderParams)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to create order: %v", err)
	}

	return &wrt_orders_v1.CreateOrderResponse{
		Status: "success",
	}, nil
}

func (h *OrdersGrpcHandler) GetCustomerOrders(ctx context.Context, req *wrt_orders_v1.GetCustomerOrdersRequest) (*wrt_orders_v1.GetCustomerOrdersResponse, error) {
	if req.GetCustomerID() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid customer id")
	}

	domainOrders, err := h.ordersService.GetCustomerOrders(ctx, req.GetCustomerID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to get orders: %v", err)
	}

	orders := []*wrt_orders_v1.Order{}
	for _, v := range domainOrders {
		orders = append(orders, fromDomainOrder(v))
	}

	return &wrt_orders_v1.GetCustomerOrdersResponse{
		Orders: orders,
	}, nil
}
