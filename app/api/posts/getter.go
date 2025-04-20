package posts

import (
	"database/sql"
	"fmt"
	"strings"

	"forum/app/handlers"
	"forum/app/modules"
	"forum/app/modules/errors"
	"forum/app/modules/log"
)

func generateBulkPostsQuery(categories []string, lastId string) (sqlQuery string, params []any) {
	sqlQuery = "SELECT DISTINCT p.id,user_id,title,content,created_at FROM posts p "
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
		sqlQuery += "p.id < ? "
	}

	sqlQuery += "ORDER BY p.id DESC LIMIT 10;"
	return
}

func GetSinglePost(conn *modules.Connection, forumDB *sql.DB) {
	conn.GetUser(forumDB)
	postId := conn.Path[2]
	sqlQuery := "SELECT id,user_id,title,content,created_at FROM posts WHERE id=?"
	posts, err := FetchPosts(sqlQuery, []any{postId}, conn.User.Id, forumDB)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}
	if len(posts) == 0 {
		conn.Error(errors.BadRequestError("invalid post id"))
		return
	}
	conn.Respond(posts[0])
}

func GetBulkPosts(conn *modules.Connection, forumDB *sql.DB) {
	if !conn.IsAuthenticated(forumDB) {
		return
	}
	queries := conn.Req.URL.Query()
	categories := queries["category"]
	lastId := queries.Get("lastId")
	conn.GetUser(forumDB)
	sqlQuery, params := generateBulkPostsQuery(categories, lastId)
	posts, err := FetchPosts(sqlQuery, params, conn.User.Id, forumDB)
	if err != nil {
		conn.Error(errors.HttpInternalServerError)
		return
	}

	conn.Respond(posts)
}

func FetchPosts(sqlQuery string, params []any, userId int, forumDB *sql.DB) (posts []modules.Post, err error) {
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
			continue
		}

		post.Likes, post.Dislikes, post.Reaction = handlers.GetReactions(post.Id, 1, userId, forumDB)
		err = post.Publisher.GetPublicUser(forumDB)
		if err != nil {
			if err == sql.ErrNoRows {
				post.Publisher.Username = "deleted user"
			} else {
				log.Error(err)
				continue
			}
		}
		post.Content.Categories, err = getPostCategories(post.Id, forumDB)
		if err != nil {
			log.Error(err)
			continue
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

func getPostCategories(postId string, forumDB *sql.DB) (categories []string, err error) {
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
