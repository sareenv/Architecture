package Errors

import "errors"

// Fixed errors just for the user status updates.
var (
	ErrUserStatusDowngrade = errors.New("user cannot be downgraded")
	ErrUserStatusUpgrade   = errors.New("user cannot be upgraded")
	ErrUserStatusNetwork   = errors.New("network error while changing the user status")
	ErrUseStatusUnknown    = errors.New("user status cannot be changed unknown status provided")
)
