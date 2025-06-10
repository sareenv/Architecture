package Models

import "time"

type User struct {
	// should be uuid, but ignore just for prototype.
	UserID    string
	FullName  string
	Email     string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id string, fullName string, email string, status Status) *User {
	return &User{
		UserID:    id,
		FullName:  fullName,
		Email:     email,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
