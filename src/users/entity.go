package users

import "time"

type User struct {
	ID          int
	Name        string
	Address     string
	PhoneNumber string
	Email       string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
