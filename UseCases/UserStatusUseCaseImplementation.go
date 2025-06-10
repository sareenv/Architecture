package UseCases

import (
	"Architecture/Errors"
	"Architecture/Models"
)

type UserStatusUseCaseImplementation struct {
	user *Models.User
}

// CheckUserStatusChange A helper method required to check if the status can be changed.
func (u UserStatusUseCaseImplementation) CheckUserStatusChange(currentStatus Models.Status,
	newStatus Models.Status, operation Models.StatusChangeOperation) (bool, error) {
	switch operation {
	case Models.Upgrade:
		if newStatus <= currentStatus || newStatus == Models.Premium {
			return false, Errors.ErrUserStatusUpgrade
		}
	case Models.Downgrade:
		if (newStatus > currentStatus) || newStatus == Models.Freemium {
			return false, Errors.ErrUserStatusDowngrade
		}
	default:
		return false, Errors.ErrUseStatusUnknown
	}
	return true, nil
}

func (u UserStatusUseCaseImplementation) UpgradeUserStatus(newStatus Models.Status) (bool, error) {
	// maybe a repository could be invoked here too or kept it separate.
	return u.CheckUserStatusChange(u.user.Status, newStatus, Models.Upgrade)
}

func (u UserStatusUseCaseImplementation) DowngradeUserStatus(newStatus Models.Status) (bool, error) {
	return u.CheckUserStatusChange(u.user.Status, newStatus, Models.Upgrade)
}

func NewUserStatusUseCaseImplementation(user *Models.User) UserStatusUseCase {
	return UserStatusUseCaseImplementation{
		user: user,
	}
}
