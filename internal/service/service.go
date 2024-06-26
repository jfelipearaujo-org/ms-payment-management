package service

import (
	"context"

	"github.com/jfelipearaujo-org/ms-payment-management/internal/entity/payment_entity"
)

type CreatePaymentService[T any] interface {
	Handle(ctx context.Context, request T) (*payment_entity.Payment, error)
}

type GetPaymentByIDService[T any] interface {
	Handle(ctx context.Context, request T) (payment_entity.Payment, error)
}

type GetPaymentsByOrderIDService[T any] interface {
	Handle(ctx context.Context, request T) ([]payment_entity.Payment, error)
}

type UpdatePaymentService[T any] interface {
	Handle(ctx context.Context, request T) (*payment_entity.Payment, error)
}

// ---

type CreatePaymentGatewayService[T any] interface {
	Handle(ctx context.Context, request T) error
}
