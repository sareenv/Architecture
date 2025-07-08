// PaymentUseCase_test.go
package Tests

import (
	"Architecture/Models"
	"Architecture/UseCases"
	"regexp"
	"testing"
	"time"
)

// TestCheckPaymentExpiryDate validates the CheckPaymentExpiryDate method by testing various card expiry date scenarios.
func TestCheckPaymentExpiryDate(t *testing.T) {
	useCase := &UseCases.PaymentMethodValidatorUseCaseImplementation{}
	currentTime := time.Now()
	currentYear := currentTime.Year() % 100
	currentMonth := int(currentTime.Month())

	tests := []struct {
		name     string
		input    Models.PaymentMethodInfo
		expected bool
	}{
		{"validExpiryDate", Models.PaymentMethodInfo{CardExpiryDate: "12/30"}, true},
		{"expiredCard", Models.PaymentMethodInfo{CardExpiryDate: "01/20"}, false},
		{"invalidMonth", Models.PaymentMethodInfo{CardExpiryDate: "13/30"}, false},
		{"invalidFormat", Models.PaymentMethodInfo{CardExpiryDate: "2023-12"}, false},
		{"emptyExpiryDate", Models.PaymentMethodInfo{CardExpiryDate: ""}, false},
		{"sameMonthAndYear", Models.PaymentMethodInfo{
			CardExpiryDate: formatExpiry(currentMonth, currentYear)}, true},
		{"pastSameYear", Models.PaymentMethodInfo{
			CardExpiryDate: formatExpiry(currentMonth-1, currentYear)}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := useCase.CheckPaymentExpiryDate(test.input)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// TestCheckCardType tests the functionality of the CheckCardType method for various valid and invalid card types and numbers.
func TestCheckCardType(t *testing.T) {
	useCase := &UseCases.PaymentMethodValidatorUseCaseImplementation{}

	tests := []struct {
		name     string
		input    Models.PaymentMethodInfo
		expected bool
	}{
		{"validVisa", Models.PaymentMethodInfo{CardType: "VISA", CardNumber: "4111111111111111"}, true},
		{"validMastercard", Models.PaymentMethodInfo{CardType: "MASTERCARD", CardNumber: "5105105105105100"}, true},
		{"validAmex", Models.PaymentMethodInfo{CardType: "AMEX", CardNumber: "371449635398431"}, true},
		{"invalidCardType", Models.PaymentMethodInfo{CardType: "UNKNOWN", CardNumber: "123456"}, false},
		{"invalidVisaNumber", Models.PaymentMethodInfo{CardType: "VISA", CardNumber: "411111"}, false},
		{"invalidMastercardNumber", Models.PaymentMethodInfo{CardType: "MASTERCARD", CardNumber: "1234567890123456"}, false},
		{"invalidAmexNumber", Models.PaymentMethodInfo{CardType: "AMEX", CardNumber: "1234567890123"}, false},
		{"emptyCardType", Models.PaymentMethodInfo{CardType: "", CardNumber: "4111111111111111"}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := useCase.CheckCardType(test.input)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// TestValidateCVVNumber tests the ValidateCVVNumber function of the PaymentMethodValidatorUseCaseImplementation struct.
func TestValidateCVVNumber(t *testing.T) {
	useCase := &UseCases.PaymentMethodValidatorUseCaseImplementation{}

	tests := []struct {
		name     string
		input    Models.PaymentMethodInfo
		expected bool
	}{
		{"validVisaCVV", Models.PaymentMethodInfo{CardType: "VISA", CardCVV: "123"}, true},
		{"validMastercardCVV", Models.PaymentMethodInfo{CardType: "MASTERCARD", CardCVV: "456"}, true},
		{"validAmexCVV", Models.PaymentMethodInfo{CardType: "AMEX", CardCVV: "1234"}, true},
		{"invalidVisaCVV", Models.PaymentMethodInfo{CardType: "VISA", CardCVV: "12"}, false},
		{"invalidAmexCVV", Models.PaymentMethodInfo{CardType: "AMEX", CardCVV: "123"}, false},
		{"emptyCVV", Models.PaymentMethodInfo{CardType: "VISA", CardCVV: ""}, false},
		{"nonNumericCVV", Models.PaymentMethodInfo{CardType: "VISA", CardCVV: "abc"}, false},
		{"longVisaCVV", Models.PaymentMethodInfo{CardType: "VISA", CardCVV: "12345"}, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := useCase.ValidateCVVNumber(test.input)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// formatExpiry generates a formatted MM/YY string for a given month and year, adjusting for out-of-range month values.
func formatExpiry(month, year int) string {
	if month < 1 {
		month += 12
		year--
	}
	return regexp.MustCompile(`^\d{1,2}/\d{2}$`).ReplaceAllString(
		time.Date(
			2000+year, time.Month(month), 1, 0, 0, 0, 0, time.UTC,
		).Format("01/06(2006)"),
		"Invalid",
	)
}
