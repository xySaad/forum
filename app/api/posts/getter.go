package posts

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"forum/app/modules"
)

func GetPosts(conn *modules.Connection, forumDB *sql.DB) error {
	var posts []modules.Post
	qreuries := conn.Req.URL.Query()
	if err := ValidQueries(qreuries); err != nil {
		return err
	}

	categories := qreuries["categories"][0]
	page, err := strconv.Atoi(qreuries["page"][0])
	if err != nil {
		return errors.New("invalide page")
	}

	err = fetchPosts(&posts, categories, page, forumDB)
	if err == nil {
		conn.Respond(posts)
	}
	return err
}

func ValidQueries(queries url.Values) error {
	if _, exits := queries["categories"]; !exits {
		return errors.New("categories missing")
	}
	if _, exits := queries["page"]; !exits {
		return errors.New("page missing")
	}
	return nil
}

func fetchPosts(posts *[]modules.Post, categories string, page int, forumDB *sql.DB) error {
	const limit = 10
	offset := (page - 1) * limit
	query := GenerateQuery(categories)
	rows, err := forumDB.Query(query, limit, offset)
	if err != nil {
		return fmt.Errorf("error querying all posts: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post modules.Post
		if err := rows.Scan(&post.Publisher.Id, &post.ID, &post.Title, &post.Text, &post.Categories, &post.CreationTime); err != nil {
			return fmt.Errorf("error scanning post: %v", err)
		}
		err := GetPublicUser(&post.Publisher, forumDB)
		if err != nil {
			return err
		}
		*posts = append(*posts, post)
	}

	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func GenerateQuery(categories string) string {
	query := `SELECT user_id, id, title, content, categories, created_at 
              FROM posts
			  `
	first := true
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

func GetCategoriesFromMask(mask string) []string {
	categoryMap := map[string]int{
		"Sport":      0,
		"Technology": 1,
		"Finance":    2,
		"Science":    3,
	}

	categories := []string{}
	for i, c := range mask {
		if c == '1' {
			for category, idx := range categoryMap {
				if idx == i {
					categories = append(categories, category)
					break
				}
			}
		}
	}

	return categories
}
