package handler

import (
	wrt_orders_v1 "github.com/wrtgvr/go-food-order-ms/services/common/genproto/orders"
	"github.com/wrtgvr/go-food-order-ms/services/orders/domain"
)

func fromDomainOrder(order *domain.Order) *wrt_orders_v1.Order {
	return &wrt_orders_v1.Order{
		OrderID:    order.OrderID,
		CustomerID: order.CustomerID,
		ProductID:  order.ProductID,
		Quantity:   order.Quantity,
	}
}

func toDomainCreateOrder(req *wrt_orders_v1.CreateOrderRequest) *domain.CreateOrderParams {
	return &domain.CreateOrderParams{
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}
}
