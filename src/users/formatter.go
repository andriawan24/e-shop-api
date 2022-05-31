package users

type UserFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	AccessToken string `json:"access_token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{}
	formatter.ID = user.ID
	formatter.Name = user.Name
	formatter.Email = user.Email
	formatter.PhoneNumber = user.PhoneNumber
	formatter.Address = user.Address
	formatter.AccessToken = token

	return formatter
}
