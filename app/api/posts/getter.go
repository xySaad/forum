package posts

import (
	"database/sql"
	"fmt"
	"strings"

	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

func GenerateGetPostSqlQuery(categories []string, lastId string) (sqlQuery string, params []any) {
	sqlQuery = "SELECT p.id,user_id,title,content,created_at FROM posts p "
	if categories != nil {
		var placeholders []string
		for _, category := range categories {
			if category != "" {
				params = append(params, category)
				placeholders = append(placeholders, "?")
			}
		}
		if placeholders != nil {
			sqlQuery += "JOIN post_categories pc ON p.id=pc.post_id JOIN categories c ON c.id=pc.category_id WHERE c.name in"
			sqlQuery += fmt.Sprintf(" (%v) ", strings.Join(placeholders, ","))
		}
	}

	if lastId != "" {
		params = append(params, lastId)
		if categories != nil {
			sqlQuery += "AND "
		} else {
			sqlQuery += "WHERE "
		}
		sqlQuery += "p.created_at > (SELECT created_at FROM posts WHERE id = ?) "
	}

	sqlQuery += "ORDER BY p.created_at DESC LIMIT 10;"
	return
}

func GetPosts(conn *modules.Connection, forumDB *sql.DB) {
	queries := conn.Req.URL.Query()
	categories := queries["category"]
	lastId := queries.Get("lastId")

	sqlQuery, params := GenerateGetPostSqlQuery(categories, lastId)
	posts, err := fetchPosts(sqlQuery, params, conn.UserId, forumDB)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}

	conn.Respond(posts)
}

func fetchPosts(sqlQuery string, params []any, user_id int, forumDB *sql.DB) (posts []modules.Post, err error) {
	rows, err := forumDB.Query(sqlQuery, params...)
	if err != nil {
		log.Error(sqlQuery, params, err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var post modules.Post
		err = rows.Scan(&post.Id, &post.Publisher.Id,
			&post.Content.Title, &post.Content.Text, &post.CreationTime)
		if err != nil {
			log.Error(err)
			return
		}
		post.Publisher.GetPublicUser(forumDB)
		// post.Likes, post.Dislikes, post.Reaction = handlers.GetReactions(post.Id, 1, user_id, forumDB)
		// err = post.Publisher.GetPublicUser(forumDB)
		// if err != nil {
		// 	log.Error(err)
		// 	return
		// }
		post.Content.Categories, err = GetPostCategories(post.Id, forumDB)
		if err != nil {
			log.Error(err)
			return
		}
		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		log.Error(err)
		return
	}
	return
}

func GetPostCategories(postId int, forumDB *sql.DB) (categories []string, err error) {
	categories = make([]string, 4)
	sqlQuery := `
        SELECT categories.name 
        FROM post_categories 
        INNER JOIN categories ON post_categories.category_id = categories.id
        WHERE post_categories.post_id = ?
    `

	rows, err := forumDB.Query(sqlQuery, postId)
	if err != nil {
		return
	}
	defer rows.Close()

	for i := 0; rows.Next() && i < len(categories); i++ {
		err = rows.Scan(&categories[i])
		if err != nil {
			return
		}
	}
	err = rows.Err()
	return
}
