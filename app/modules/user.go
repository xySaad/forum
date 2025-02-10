package modules

import "database/sql"

type User struct {
	Username       string  `json:"username"`
	Id             int     `json:"id"`
	ProfilePicture *string `json:"profilePicture"`
}

func (u *User) GetPublicUser(db *sql.DB) (err error) {
	qreury := `SELECT id,username,profile_picture FROM users WHERE id=?`
	err = db.QueryRow(qreury, u.Id).Scan(&u.Id, &u.Username, &u.ProfilePicture)
	return
}
