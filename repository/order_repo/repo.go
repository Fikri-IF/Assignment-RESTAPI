package order_repo

import "kominfo-assignment-2/entity"

type Repository interface {
	CreateOrderWithItems(order entity.Order, items []entity.Item) error
	GetOrders() ([]entity.Order, error)
	UpdateOrderWithItems(order entity.Order, items []entity.Item, orderId string) error
	DeleteOrder(orderId string) error
}
