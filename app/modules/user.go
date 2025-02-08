package modules

type User struct {
	Username       string  `json:"username"`
	Id             int `json:"id"`
	ProfilePicture *string `json:"profilePicture"`
}
