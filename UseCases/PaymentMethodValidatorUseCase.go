package UseCases

import (
	"Architecture/Models"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type PaymentMethodValidatorUseCase interface {
	Execute(info Models.PaymentMethodInfo) (bool, error)
}

type PaymentMethodValidatorUseCaseImplementation struct{}

func NewPaymentMethodValidatorUseCase() PaymentMethodValidatorUseCase {
	return &PaymentMethodValidatorUseCaseImplementation{}
}

func (p *PaymentMethodValidatorUseCaseImplementation) Execute(info Models.PaymentMethodInfo) (bool, error) {
	// Check all validation rules
	if !p.CheckCardType(info) {
		return false, fmt.Errorf("invalid card type or number")
	}

	if !p.CheckPaymentExpiryDate(info) {
		return false, fmt.Errorf("card is expired or has invalid expiry date")
	}

	if !p.ValidateCVVNumber(info) {
		return false, fmt.Errorf("invalid CVV number")
	}

	return true, nil
}

func (p *PaymentMethodValidatorUseCaseImplementation) CheckPaymentExpiryDate(info Models.PaymentMethodInfo) bool {
	// Parse the expiry date (expected format: "MM/YY")
	parts := strings.Split(info.CardExpiryDate, "/")
	if len(parts) != 2 {
		return false
	}

	currentTime := time.Now()
	currentYear := currentTime.Year() % 100 // Get last two digits of year
	currentMonth := int(currentTime.Month())

	// Parse month and year from card
	var month, year int
	_, err := fmt.Sscanf(parts[0], "%d", &month)
	if err != nil || month < 1 || month > 12 {
		return false
	}
	_, err = fmt.Sscanf(parts[1], "%d", &year)
	if err != nil {
		return false
	}

	// Check if card is expired
	if year < currentYear || (year == currentYear && month < currentMonth) {
		return false
	}

	return true
}

func (p *PaymentMethodValidatorUseCaseImplementation) CheckCardType(info Models.PaymentMethodInfo) bool {
	// Common card types validation
	cardPatterns := map[string]string{
		"VISA":       "^4[0-9]{12}(?:[0-9]{3})?$",
		"MASTERCARD": "^5[1-5][0-9]{14}$",
		"AMEX":       "^3[47][0-9]{13}$",
	}

	pattern, exists := cardPatterns[strings.ToUpper(info.CardType)]
	if !exists {
		return false
	}

	matched, err := regexp.MatchString(pattern, info.CardNumber)
	if err != nil {
		return false
	}

	return matched
}

func (p *PaymentMethodValidatorUseCaseImplementation) ValidateCVVNumber(info Models.PaymentMethodInfo) bool {
	// CVV validation rules
	cvvLength := len(info.CardCVV)

	// AMEX cards have 4-digit CVV, others have 3-digit CVV
	if strings.ToUpper(info.CardType) == "AMEX" {
		return cvvLength == 4 && regexp.MustCompile(`^\d{4}$`).MatchString(info.CardCVV)
	}

	return cvvLength == 3 && regexp.MustCompile(`^\d{3}$`).MatchString(info.CardCVV)
}
