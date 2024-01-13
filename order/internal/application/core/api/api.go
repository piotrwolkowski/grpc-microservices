package api

import (
	"strings"

	"github.com/piotrwolkowski/grpc-microservices/order/internal/application/core/domain"
	"github.com/piotrwolkowski/grpc-microservices/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db       ports.DBPort
	payments ports.PaymentPort
}

func NewApplication(db ports.DBPort, payments ports.PaymentPort) *Application {
	return &Application{
		db:       db,
		payments: payments,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	err = a.payments.Charge(&order)
	if err != nil {
		st := status.Convert(err)
		var allErrors []string
		for _, detail := range st.Details() {
			switch t := detail.(type) {
			case *errdetails.BadRequest:
				for _, violation := range t.GetFieldViolations() {
					allErrors = append(allErrors, violation.Description)
				}
			}
		}
		fieldErr := &errdetails.BadRequest_FieldViolation{
			Field:       "payment",
			Description: strings.Join(allErrors, "\n"),
		}
		badReq := &errdetails.BadRequest{}
		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
		orderStatus := status.New(codes.InvalidArgument, "order creation failed - payment issue")
		statusWithDetails, _ := orderStatus.WithDetails(badReq)

		return domain.Order{}, statusWithDetails.Err()
	}

	return order, nil
}

// Example of an alternative error handling
// func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
// 	err := a.db.Save(&order)
// 	if err != nil {
// 		return domain.Order{}, err
// 	}
// 	err = a.payments.Charge(&order)
// 	if err != nil {
// 		st, _ := status.FromError(err)
// 		fieldErr := &errdetails.BadRequest_FieldViolation{
// 			Field:       "payment",
// 			Description: st.Message(),
// 		}
// 		badReq := &errdetails.BadRequest{}
// 		badReq.FieldViolations = append(badReq.FieldViolations, fieldErr)
// 		orderStatus := status.New(codes.InvalidArgument, "order creation failed - payment issue")
// 		statusWithDetails, _ := orderStatus.WithDetails(badReq)
// 		return domain.Order{}, statusWithDetails.Err()
// 	}
// 	return order, nil
// }

func (a Application) GetOrder(id int64) (domain.Order, error) {
	// stringID := strconv.Itoa(int(id))
	order, err := a.db.Get(id)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
