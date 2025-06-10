package main

import (
	"Architecture/Models"
	"Architecture/UseCases"
	"fmt"
)

func main() {
	var newUser *Models.User = Models.NewUser(
		"DAXUI",
		"Vinayak Sareen",
		"contact@sareenv.com",
		Models.Freemium,
	)
	userStatusUseCase := UseCases.NewUserStatusUseCaseImplementation(newUser)
	isUpgraded, upgradeErr := userStatusUseCase.UpgradeUserStatus(Models.Basic)
	if upgradeErr != nil {
		fmt.Printf("Failed to upgrade user status: %v\n", upgradeErr)
		return
	}
	if isUpgraded {
		fmt.Println("User status successfully upgraded to Basic")
	} else {
		fmt.Println("User status upgrade was not processed")
	}
}
