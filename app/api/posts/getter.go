package posts

import (
	"database/sql"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

func GetPosts(conn *modules.Connection, forumDB *sql.DB) {
	qeuries := conn.Req.URL.Query()
	posts := []modules.Post{}
	categories, from, page, err := ValidQueries(qeuries)
	if err != nil {
		conn.Error(err)
		return
	}
	if categories == "" {
		posts, err = fetchPosts(page, from, forumDB)
	} else {
		posts, err = fetchPostsByCategories(strings.Split(categories, ","), from, page, forumDB)
	}
	if err == nil {
		conn.Respond(posts)
	} else {
		conn.Error(err)
	}
}

func ValidQueries(queries url.Values) (categories, from string, page int, err *errors.HttpError) {
	if _, exits := queries["categories"]; !exits {
		categories = ""
	} else {
		categories = queries["categories"][0]
	}
	if _, exits := queries["from"]; !exits {
		err = &errors.HttpError{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Code:    http.StatusBadRequest,
			Details: "mising from",
		}
		return
	}

	if _, exits := queries["page"]; !exits {
		err = &errors.HttpError{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Code:    http.StatusBadRequest,
			Details: "mising page",
		}
		return
	}
	from = queries["from"][0]
	page, Err := strconv.Atoi(queries["page"][0])
	if Err != nil || page < 0 {
		err = &errors.HttpError{
			Status:  http.StatusBadRequest,
			Message: "bad request",
			Code:    http.StatusBadRequest,
			Details: "page should be a positive number",
		}
		return
	}
	if nFrom, Err := strconv.Atoi(from); Err != nil || nFrom <= 0 {
		if Err != nil || nFrom < 0 {
			err = &errors.HttpError{
				Status:  http.StatusBadRequest,
				Message: "bad request",
				Code:    http.StatusBadRequest,
				Details: "from should be a positive number",
			}
			return "", "", 0, err
		}
		from = ""
	}
	return
}

func fetchPostsByCategories(categories []string, from string, page int, db *sql.DB) ([]modules.Post, *errors.HttpError) {
	posts := []modules.Post{}
	const limit = 10
	params := make([]interface{}, len(categories))
	placeholders := []string{}
	for i, category := range categories {
		params[i] = category
		placeholders = append(placeholders, "?")
	}
	query := `SELECT DISTINCT user_id, id, title, content, created_at FROM posts
INNER JOIN posts_categories on posts_categories.post_id = posts.post_id
INNER JOIN categories on posts_categories.category_id = categories.id
WHERE category.name IN (` + strings.Join(placeholders, ", ") + `)`
	if from != "" {
		query += ` AND posts.id <= ` + from
	}
	query += `
ORDER BY posts.created_at DESC LIMIT ? OFFSET ?`
	params = append(params, limit, page*limit)
	rows, err := db.Query(query, params...)
	herr := &errors.HttpError{}
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.HttpInternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var post modules.Post
		if err := rows.Scan(&post.Publisher.Id, &post.ID, &post.Title, &post.Text, &post.CreationTime); err != nil {
			return nil, errors.HttpInternalServerError
		}

		err := GetPublicUser(&post.Publisher, db)
		if err != nil {
			log.Warn(err)
		}
		herr = GetPostCategories(&post.Categories, post.ID, db)
		if err != nil {
			log.Warn(err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.HttpInternalServerError
	}
	return posts, herr
}

func fetchPosts(page int, from string, forumDB *sql.DB) (posts []modules.Post, herr *errors.HttpError) {
	const limit = 10
	offset := (page - 1) * limit
	query := `SELECT  user_id, id, title, content, created_at FROM posts`
	if from != "" {
		query += ` WHERE id <= ` + from
	}
	query += ` ORDER BY created_at DESC OFFSET ? LIMIT ?`
	rows, err := forumDB.Query(query, offset, limit)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.HttpInternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var post modules.Post
		if err := rows.Scan(&post.Publisher.Id, &post.ID, &post.Title, &post.Text, &post.CreationTime); err != nil {
			return nil, errors.HttpInternalServerError
		}

		err := GetPublicUser(&post.Publisher, forumDB)
		if err != nil {
			log.Warn(err)
		}
		herr = GetPostCategories(&post.Categories, post.ID, forumDB)
		if err != nil {
			log.Warn(err)
		}
		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.HttpInternalServerError
	}
	return posts, nil
}

func GetPostCategories(categories *[]string, postID string, db *sql.DB) *errors.HttpError {
	query := `SELECT category_id FROM post_categories where post_id =`
	res, err := db.Query(query, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			return  nil
		}
		return  errors.HttpInternalServerError
	}
	defer res.Close()
	for res.Next() {
		id := ""
		err = res.Scan(id)
		if err != nil {
			log.Warn(err)
		}
		category := ""
		err = db.QueryRow(`SELECT name FROM categories WHERE id = ?`, id).Scan(&category)
		if err != nil {
			log.Warn(err)
		} else {
			*categories = append(*categories, category)
		}
	}
	if err := res.Err(); err != nil {
		return errors.HttpInternalServerError
	}
	return nil
}

func GetPublicUser(user *modules.User, db *sql.DB) error {
	qreury := `SELECT username,profile FROM users WHERE id=?`
	return db.QueryRow(qreury, user.Id).Scan(&user.Username, &user.ProfilePicture)
}
