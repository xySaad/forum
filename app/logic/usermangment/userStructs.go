package usermangment

type User struct {
	Username string `json:"Username"`
	 Email string	`json:"Email"`
	 Password string	`json:"Password"`
	 PasswordConfirm string	` json:"ConfirmPassword"`
}
