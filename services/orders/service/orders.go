package service

import (
	"context"

	"github.com/wrtgvr/go-food-order-ms/services/orders/domain"
)

type OrdersService struct {
	// repo
}

func NewOrdersService() *OrdersService {
	return &OrdersService{}
}

// !! Temporary storage
var ordersMap = make(map[uint32]*domain.Order)
var nextOrderId uint32 = 0

func (s *OrdersService) CreateOrder(ctx context.Context, params *domain.CreateOrderParams) error {
	ordersMap[nextOrderId] = &domain.Order{
		OrderID:    nextOrderId,
		CustomerID: params.CustomerID,
		ProductID:  params.ProductID,
		Quantity:   params.Quantity,
		Status:     domain.StatusOrderInQueue,
	}
	nextOrderId++

	return nil
}

func (s *OrdersService) GetCustomerOrders(ctx context.Context, customerId uint32) ([]*domain.Order, error) {
	orders := []*domain.Order{}

	for _, order := range ordersMap {
		if order.CustomerID == customerId {
			orders = append(orders, order)
		}
	}

	return orders, nil
}
