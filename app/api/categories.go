package api

import (
	"database/sql"

	"forum/app/modules"
	"forum/app/modules/errors"
)

func GetAllCategories(conn *modules.Connection, forumDB *sql.DB) {
	categories := []string{}
	rows, err := forumDB.Query("SELECT name from categories")
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}
	for rows.Next() {
		var category string
		err = rows.Scan(&category)
		if err != nil {
			conn.Error(errors.HttpInternalServerError)
			return
		}
		categories = append(categories, category)
	}
	conn.Respond(categories)
}
