package modules

type User struct 
{
	Username       string 	`json:"username"`
	Id             string
	ProfilePicture any 		`json:"profilePicture"`
}
