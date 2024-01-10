package main

import (
	"log"

	"github.com/piotrwolkowski/microservices/order/config"
	"github.com/piotrwolkowski/microservices/order/internal/adapters/db"
	"github.com/piotrwolkowski/microservices/order/internal/adapters/grpc"
	"github.com/piotrwolkowski/microservices/order/internal/adapters/payment"
	"github.com/piotrwolkowski/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to create db adapter: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
