package ports

import "github.com/piotrwolkowski/microservices/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(order *domain.Order) error
}
