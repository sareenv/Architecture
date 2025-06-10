package UseCases

import "Architecture/Models"

// UserStatusUseCase contains the business logic here
type UserStatusUseCase interface {
	// UpgradeUserStatus we can only upgrade the status.
	UpgradeUserStatus(newStatus Models.Status) (bool, error)
	// DowngradeUserStatus we can only downgrade the status.
	DowngradeUserStatus(newStatus Models.Status) (bool, error)
}
