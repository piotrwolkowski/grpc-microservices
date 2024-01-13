package ports

import "github.com/piotrwolkowski/grpc-microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}
