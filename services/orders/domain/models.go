package domain

type CreateOrderParams struct {
	CustomerID uint32
	ProductID  uint32
	Quantity   uint32
}

type Order struct {
	OrderID    uint32
	CustomerID uint32
	ProductID  uint32
	Quantity   uint32
	Status     string
}
