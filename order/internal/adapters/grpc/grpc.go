package grpc

import (
	"context"

	"github.com/piotrwolkowski/microservices/order/internal/application/core/domain"
	"github.com/piotrwolkowski/microservices/proto/golang/order"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, item := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
		})
	}

	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{OrderId: result.ID}, nil
}

func (a Adapter) Get(ctx context.Context, request *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	orderFromDb, err := a.api.GetOrder(request.OrderId)
	if err != nil {
		return nil, err
	}

	var orderItems []*order.OrderItem
	for _, item := range orderFromDb.OrderItems {
		orderItems = append(orderItems, &order.OrderItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
		})
	}

	return &order.GetOrderResponse{
		UserId:     orderFromDb.CustomerID,
		OrderItems: orderItems,
	}, nil
}
