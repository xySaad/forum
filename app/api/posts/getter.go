package posts

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"forum/app/modules"
	"forum/app/modules/log"
)

func GetPosts(conn *modules.Connection, forumDB *sql.DB) error {
	qeuries := conn.Req.URL.Query()
	categories, from, page, err := ValidQueries(qeuries)
	if err != nil {
		return err
	}
	posts, err := fetchPosts(categories, page, from, forumDB)
	if err == nil {
		conn.Respond(posts)
	}
	return err
}

func ValidQueries(queries url.Values) (categories, from string, page int, err error) {
	if _, exits := queries["categories"]; !exits {
		err = errors.New("categories missing")
		return
	}
	if _, exits := queries["from"]; !exits {
		err = errors.New("from missing")
		return
	}

	if _, exits := queries["page"]; !exits {
		err = errors.New("page missing")
		return
	}
	from = queries["from"][0]
	categories = queries["categories"][0]
	if categories == "" || len(categories) != 4 {
		err = errors.New("invalid categories")
		return
	}
	for _, v := range categories {
		if !(v == '0' || v == '1') {
			err = errors.New("invalid categories")
			return
		}
	}
	page, err = strconv.Atoi(queries["page"][0])
	if err != nil || page < 0 {
		err = errors.New("page should be greater than 0")
		return
	}
	if nFrom, err := strconv.Atoi(from); err != nil || nFrom <= 0 {
		if err != nil {
			err = errors.New("from should be a positive number")
			return "", "", 0, err
		}
		from = ""
	}

	return
}

func fetchPosts(categories string, page int, from string, forumDB *sql.DB) (posts []modules.Post, err error) {
	const limit = 10
	offset := (page - 1) * limit
	query := GenerateQuery(categories, from)
	fmt.Println(query)
	rows, err := forumDB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error querying all posts: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post modules.Post
		if err := rows.Scan(&post.Publisher.Id, &post.ID, &post.Title, &post.Text, &post.Categories, &post.CreationTime); err != nil {
			return nil, fmt.Errorf("error scanning post: %v", err)
		}
		err := GetPublicUser(&post.Publisher, forumDB)
		if err != nil {
			log.Warn(err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func GenerateQuery(categories string, from string) string {
	query := `SELECT user_id, id, title, content, categories, created_at 
              FROM posts `
	first := true
	if from != "" {
		query += ` WHERE id<=` + from + ` `
		first = false
	}
	for i, v := range categories {
		if v == '1' {
			if first {
				query += `WHERE `
				first = false
			} else {
				query += `AND `
			}
			idx := strconv.Itoa(i)
			query += `SUBSTRING(categories,` + idx + `,1)='1' `
		}
	}
	query += `ORDER BY created_at DESC 
	LIMIT ? OFFSET ?`
	return query
}

func GetPublicUser(user *modules.User, db *sql.DB) error {
	qreury := `SELECT username,profile FROM users WHERE id=?`
	return db.QueryRow(qreury, user.Id).Scan(&user.Username, &user.ProfilePicture)
}
