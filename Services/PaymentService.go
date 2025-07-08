package Services

import "Architecture/UseCases"

type PaymentServiceDependency struct {
	UseCases.PaymentMethodValidatorUseCase
	UseCases.ProcessPaymentUseCase
	// manage a payment method use case.
}

type PaymentService interface {
	PaymentServiceDependency
}
