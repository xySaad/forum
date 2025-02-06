package modules

type User struct {
	Username       string  `json:"username"`
	Id             string  `json:"id"`
	ProfilePicture *string `json:"profilePicture"`
}
