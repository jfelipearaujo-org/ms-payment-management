package update

import (
	"github.com/go-playground/validator/v10"
	"github.com/jfelipearaujo-org/ms-payment-management/internal/shared/custom_error"
)

type UpdatePaymentDTO struct {
	PaymentId string `json:"payment_id" validate:"required,uuid4"`
	Approved  bool   `json:"approved"`
}

func (dto *UpdatePaymentDTO) Validate() error {
	validator := validator.New()

	if err := validator.Struct(dto); err != nil {
		return custom_error.ErrRequestNotValid
	}

	return nil
}
