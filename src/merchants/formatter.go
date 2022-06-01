package merchants

import "time"

type MerchantFormatter struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	Tagline     string    `json:"tagline"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormatMerchant(merchant Merchant) MerchantFormatter {
	formatter := MerchantFormatter{}
	formatter.ID = merchant.ID
	formatter.Name = merchant.Name
	formatter.Address = merchant.Address
	formatter.PhoneNumber = merchant.PhoneNumber
	formatter.Tagline = merchant.Tagline
	formatter.Description = merchant.Description
	formatter.CreatedAt = merchant.CreatedAt
	formatter.UpdatedAt = merchant.UpdatedAt

	return formatter
}
