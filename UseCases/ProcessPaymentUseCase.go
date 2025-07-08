package UseCases

import (
	"Architecture/Models"
	"errors"
	"fmt"
)

type ProcessPaymentUseCase interface {
	Execute(paymentInfo Models.PaymentMethodInfo, amount Models.Payment) (bool, error)
}

type ProcessPaymentUseCaseImpl struct {
	paymentValidator PaymentMethodValidatorUseCase
}

func NewProcessPaymentUseCase(validator PaymentMethodValidatorUseCase) ProcessPaymentUseCase {
	return &ProcessPaymentUseCaseImpl{
		paymentValidator: validator,
	}
}

func (p *ProcessPaymentUseCaseImpl) Execute(paymentInfo Models.PaymentMethodInfo, amount Models.Payment) (bool, error) {
	if amount.Amount <= 0 {
		return false, errors.New("payment amount must be greater than zero")
	}

	if amount.Currency == "" {
		return false, errors.New("currency must be specified")
	}

	// Validate payment method first
	valid, err := p.paymentValidator.Execute(paymentInfo)
	if err != nil {
		return false, fmt.Errorf("payment validation failed: %w", err)
	}
	if !valid {
		return false, errors.New("invalid payment method")
	}

	// Process payment with an external payment provider
	success, err := p.processPaymentWithProvider(paymentInfo, amount)
	if err != nil {
		return false, fmt.Errorf("payment processing failed: %w", err)
	}

	return success, nil
}

func (p *ProcessPaymentUseCaseImpl) processPaymentWithProvider(info Models.PaymentMethodInfo, amount Models.Payment) (bool, error) {
	// This would typically integrate with a payment gateway/provider
	// Mocked for demonstration
	if info.PaymentMethod == "CARD" || info.PaymentMethod == "CREDIT_CARD" {
		// Simulate successful card payment
		return true, nil
	}
	// maybe like a btc or etc which is not
	return false, errors.New("unsupported payment method")
}
