package auth

type response struct {
	Url     string `json:"url"`
	Message string `json:"message"`
}
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
