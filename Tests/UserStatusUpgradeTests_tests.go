package Tests

import (
	"Architecture/Models"
	"Architecture/UseCases"
	"testing"
)

func TestUserStatusUpgradeWhenStatusCannotDowngrade(t *testing.T) {
	// Given a user and use case.
	testUser := Models.NewUser(
		"1",
		"test user",
		"testuser@gmail.com",
		Models.Freemium,
	)
	useCase := UseCases.NewUserStatusUseCaseImplementation(testUser)
	result, _ := useCase.DowngradeUserStatus(Models.Freemium)
	if result != false {
		t.Errorf("Incorrect result")
	}
}
