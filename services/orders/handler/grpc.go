package handler

import (
	"context"

	wrt_orders_v1 "github.com/wrtgvr/go-food-order-ms/services/common/genproto/orders"
	"github.com/wrtgvr/go-food-order-ms/services/orders/domain"
	"google.golang.org/grpc"
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
	createOrderParams := &domain.CreateOrderParams{
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}
	// TODO: funcs-converters (CreateOrderRequest -> CreateOrderParams)
	// TODO: fields validation

	err := h.ordersService.CreateOrder(ctx, createOrderParams)
	if err != nil {
		return nil, err
	}

	return &wrt_orders_v1.CreateOrderResponse{
		Status: "success",
	}, nil
}

func (h *OrdersGrpcHandler) GetCustomerOrders(ctx context.Context, req *wrt_orders_v1.GetCustomerOrdersRequest) (*wrt_orders_v1.GetCustomerOrdersResponse, error) {
	domainOrders, err := h.ordersService.GetCustomerOrders(ctx, req.GetCustomerID())
	if err != nil {
		return nil, err
	}
	// TODO: fields validation

	orders := []*wrt_orders_v1.Order{}
	for _, v := range domainOrders {
		orders = append(orders, &wrt_orders_v1.Order{
			OrderID:    v.OrderID,
			CustomerID: v.CustomerID,
			ProductID:  v.ProductID,
			Quantity:   v.Quantity,
		})
	}
	// TODO: funcs-converters (domain.Order -> wrt_orders_v1.Order)

	return &wrt_orders_v1.GetCustomerOrdersResponse{
		Orders: orders,
	}, nil
}
