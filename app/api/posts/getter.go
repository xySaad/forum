package posts

import (
	"database/sql"
	"strconv"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

func GetPosts(conn *modules.Connection, forumDB *sql.DB) {
	queries := conn.Req.URL.Query()
	categories := queries["categories"]
	lastIdQuery := queries["lastId"]
	lastId := 0
	if lastIdQuery != nil {
		n, err := strconv.Atoi(lastIdQuery[0])
		if err != nil {
			conn.Error(errors.BadRequestError("invalid lastId"))
			return
		}
		lastId = n
	}
	posts := []modules.Post{}

	posts, httpErr := fetchPosts(lastId, categories, forumDB)
	if httpErr == nil {
		conn.Respond(posts)
	} else {
		conn.Error(httpErr)
	}
}

func fetchPosts(lastId int, categories []string, forumDB *sql.DB) (posts []modules.Post, httpErr *errors.HttpError) {
	sqlQuery := `SELECT 
	id,
	internal_id,
	user_internal_id,
	title, 
	content, 
	created_at 
	FROM posts `

	if lastId > 0 {
		sqlQuery += "where posts.created_at > (select created_at from posts where id = ?) "
	}

	sqlQuery += "ORDER BY posts.created_at DESC LIMIT 10;"

	if categories != nil {

	}

	rows, err := forumDB.Query(sqlQuery, lastId)
	if err != nil {
		log.Error(sqlQuery, err)
		return nil, errors.HttpInternalServerError
	}
	defer rows.Close()
	for rows.Next() {
		var internalPostId, internalUserId int
		var post modules.Post
		err = rows.Scan(&post.ID, &internalPostId, &internalUserId,
			&post.Title, &post.Text, &post.CreationTime)

		if err != nil {
			log.Error(err)
			return
		}

		post.Publisher, err = GetPublicUser(internalUserId, forumDB)
		if err != nil {
			log.Error(err)
			return
		}
		post.Categories, err = GetPostCategories(internalPostId, forumDB)
		if err != nil {
			log.Error(err)
			return
		}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		log.Error(err)
		return posts, errors.HttpInternalServerError
	}
	return posts, nil
}
func GetPostCategories(postInternalID int, forumDB *sql.DB) (categories []string, err error) {
	sqlQuery := `
        SELECT categories.name 
        FROM post_categories 
        INNER JOIN categories ON post_categories.category_id = categories.id
        WHERE post_categories.post_internal_id = ?
    `

	rows, err := forumDB.Query(sqlQuery, postInternalID)
	if err != nil {
		return
	}
	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		if i == len(categories) {
			categories = append(categories, "")
		}
		err = rows.Scan(&categories[i])
		if err != nil {
			return
		}
	}
	err = rows.Err()
	return
}

func GetPublicUser(internalUserID int, db *sql.DB) (user modules.User, err error) {
	qreury := `SELECT id, username,profile_picture FROM users WHERE internal_id=?`
	err = db.QueryRow(qreury, internalUserID).Scan(&user.Id, &user.Username, &user.ProfilePicture)
	return
}
