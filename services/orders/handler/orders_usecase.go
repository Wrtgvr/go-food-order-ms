package handler

import (
	"context"

	"github.com/wrtgvr/go-food-order-ms/services/orders/domain"
)

type OrdersService interface {
	CreateOrder(context.Context, *domain.CreateOrderParams) error
	GetOrder(context.Context, uint32) (*domain.Order, error)
}
