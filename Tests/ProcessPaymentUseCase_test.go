package Tests

import (
	"Architecture/Models"
	"Architecture/UseCases"
	"errors"
	"testing"
)

type mockValidator struct {
	isValid      bool
	errorToThrow error
}

func (m *mockValidator) Execute(info Models.PaymentMethodInfo) (bool, error) {
	if m.errorToThrow != nil {
		return false, m.errorToThrow
	}
	return m.isValid, nil
}

func TestProcessPaymentUseCaseImpl_Execute(t *testing.T) {
	tests := []struct {
		name            string
		paymentInfo     Models.PaymentMethodInfo
		amount          Models.Payment
		validatorMock   *mockValidator
		expectedSuccess bool
		expectedError   string
	}{
		{
			name: "valid card payment",
			paymentInfo: Models.PaymentMethodInfo{
				PaymentMethod: "CARD",
			},
			amount: Models.Payment{
				Amount:   100.0,
				Currency: "USD",
			},
			validatorMock:   &mockValidator{isValid: true},
			expectedSuccess: true,
			expectedError:   "",
		},
		{
			name: "invalid amount less than or equal to zero",
			paymentInfo: Models.PaymentMethodInfo{
				PaymentMethod: "CARD",
			},
			amount: Models.Payment{
				Amount:   0.0,
				Currency: "USD",
			},
			validatorMock:   &mockValidator{isValid: true},
			expectedSuccess: false,
			expectedError:   "payment amount must be greater than zero",
		},
		{
			name: "missing currency",
			paymentInfo: Models.PaymentMethodInfo{
				PaymentMethod: "CARD",
			},
			amount: Models.Payment{
				Amount:   100.0,
				Currency: "",
			},
			validatorMock:   &mockValidator{isValid: true},
			expectedSuccess: false,
			expectedError:   "currency must be specified",
		},
		{
			name: "invalid payment method",
			paymentInfo: Models.PaymentMethodInfo{
				PaymentMethod: "UNSUPPORTED",
			},
			amount: Models.Payment{
				Amount:   100.0,
				Currency: "USD",
			},
			validatorMock:   &mockValidator{isValid: true},
			expectedSuccess: false,
			expectedError:   "unsupported payment method",
		},
		{
			name: "validator fails with error",
			paymentInfo: Models.PaymentMethodInfo{
				PaymentMethod: "CARD",
			},
			amount: Models.Payment{
				Amount:   100.0,
				Currency: "USD",
			},
			validatorMock:   &mockValidator{isValid: false, errorToThrow: errors.New("validation failed")},
			expectedSuccess: false,
			expectedError:   "payment validation failed: validation failed",
		},
		{
			name: "validator returns invalid payment",
			paymentInfo: Models.PaymentMethodInfo{
				PaymentMethod: "CARD",
			},
			amount: Models.Payment{
				Amount:   100.0,
				Currency: "USD",
			},
			validatorMock:   &mockValidator{isValid: false},
			expectedSuccess: false,
			expectedError:   "invalid payment method",
		},
		{
			name: "valid credit card payment",
			paymentInfo: Models.PaymentMethodInfo{
				PaymentMethod: "CREDIT_CARD",
			},
			amount: Models.Payment{
				Amount:   200.0,
				Currency: "EUR",
			},
			validatorMock:   &mockValidator{isValid: true},
			expectedSuccess: true,
			expectedError:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := UseCases.NewProcessPaymentUseCase(tt.validatorMock)
			success, err := useCase.Execute(tt.paymentInfo, tt.amount)

			if success != tt.expectedSuccess {
				t.Errorf("expected success: %v, got: %v", tt.expectedSuccess, success)
			}

			if err != nil && err.Error() != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err.Error())
			}

			if err == nil && tt.expectedError != "" {
				t.Errorf("expected error: %v, got none", tt.expectedError)
			}
		})
	}
}
