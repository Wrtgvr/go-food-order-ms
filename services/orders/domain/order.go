package domain

import (
	"errors"
	"strings"
)

type Order struct {
	OrderID    uint32
	CustomerID uint32
	ProductID  uint32
	Quantity   uint32
	Status     string
}

const (
	StatusOrderInQueue   = "in queue"
	StatusOrderPreparing = "is preparing"
	StatusOrderReady     = "ready"
)

func (o Order) Validate() error {
	errs := make([]string, 0, 5)

	if o.OrderID <= 0 {
		errs = append(errs, "invalid order id")
	}
	if o.CustomerID <= 0 {
		errs = append(errs, "invalid customer id")
	}
	if o.ProductID <= 0 {
		errs = append(errs, "invalid product id")
	}
	if o.Quantity <= 0 {
		errs = append(errs, "invalid quantity")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}

type CreateOrderParams struct {
	CustomerID uint32
	ProductID  uint32
	Quantity   uint32
}

func (p CreateOrderParams) Validate() error {
	errs := make([]string, 0, 3)

	if p.CustomerID <= 0 {
		errs = append(errs, "invalid customer id")
	}
	if p.ProductID <= 0 {
		errs = append(errs, "invalid product id")
	}
	if p.Quantity <= 0 {
		errs = append(errs, "invalid quantity")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ", "))
	}

	return nil
}
