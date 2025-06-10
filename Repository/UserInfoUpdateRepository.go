package Repository

import (
	"Architecture/Models"
	"time"
)

type UserInfoUpdateRepository interface {

	// UpdateUserInfo updates the information of a user with the given userId and returns an update id message or an error.
	UpdateUserInfo(updatedAtTime time.Time, userId string, user Models.User) (string, error)

	// RevertUpdateInfo reverts the specified user information update identified by the given updateId.
	RevertUpdateInfo(updateId string)
}
