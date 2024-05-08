package services

import "application-design-test-master/models"

// OrderService interface defining methods for orders.
type OrderService interface {
	// CreateOrder creates a new order.
	CreateOrder(order models.Order) error
}
