package merchants

import "time"

type Merchant struct {
	ID          int
	Name        string
	Address     string
	PhoneNumber string
	Tagline     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
