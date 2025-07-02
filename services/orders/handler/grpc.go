package handler

import (
	"context"
	"log/slog"

	wrt_orders_v1 "github.com/wrtgvr/go-food-order-ms/services/common/genproto/orders"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrdersGrpcHandler struct {
	deps *HandlerDeps
	wrt_orders_v1.UnimplementedOrdersServiceServer
}

func NewOrdersGrpcHandler(grpcServer *grpc.Server, deps *HandlerDeps) {
	h := &OrdersGrpcHandler{
		deps: deps,
	}

	wrt_orders_v1.RegisterOrdersServiceServer(grpcServer, h)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *wrt_orders_v1.CreateOrderRequest) (*wrt_orders_v1.CreateOrderResponse, error) {
	h.deps.Log.Info("gRPC request",
		slog.Group("request",
			slog.Int64("Customer ID", int64(req.GetCustomerID())),
			slog.Int64("Product ID", int64(req.GetProductID())),
			slog.Int64("Quantity", int64(req.GetQuantity())),
		),
	)

	createOrderParams := toDomainCreateOrder(req)

	if err := createOrderParams.Validate(); err != nil {
		return nil, logAndWrapError(h.deps.Log, "request validation failed", err, codes.InvalidArgument)
	}

	err := h.deps.OrdersService.CreateOrder(ctx, createOrderParams)
	if err != nil {
		return nil, logAndWrapError(h.deps.Log, "unable to create order", err, codes.Internal)
	}

	return &wrt_orders_v1.CreateOrderResponse{
		Status: "success",
	}, nil
}

func (h *OrdersGrpcHandler) GetCustomerOrders(ctx context.Context, req *wrt_orders_v1.GetCustomerOrdersRequest) (*wrt_orders_v1.GetCustomerOrdersResponse, error) {
	h.deps.Log.Info("gRPC request",
		slog.Group("request",
			slog.Int64("Customer ID", int64(req.GetCustomerID())),
		),
	)

	if req.GetCustomerID() <= 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid customer id")
	}

	domainOrders, err := h.deps.OrdersService.GetCustomerOrders(ctx, req.GetCustomerID())
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
