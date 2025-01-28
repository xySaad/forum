package posts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"forum/app/modules"
)

func GetPosts(conn *modules.Connection, forumDB *sql.DB) error {
	var posts []modules.Post
	var categories []string
	var err error
	pageStr := ""

	if len(strings.Split(conn.Req.URL.Path, "/")) == 4 {
		pageStr = strings.Split(conn.Req.URL.Path, "/")[3]
	}
	page := 1

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if strings.Contains(conn.Req.URL.Path, "categories") {
		cat := strings.Split(conn.Req.URL.Path, "/categories=")[1]
		categories = strings.Split(cat, "&")

	}

	if len(categories) > 0 {
		err = fetchPostsByCategories(categories, &posts, page, forumDB)
	} else {
		err = fetchAllPosts(&posts, page, forumDB)
	}

	if err != nil {
		return fmt.Errorf("error fetching posts: %v", err)
	}

	postJSON, err := json.Marshal(posts)
	if err != nil {
		return fmt.Errorf("error marshaling posts: %v", err)
	}
	conn.Resp.Write(postJSON)

	return nil
}

func fetchPostsByCategories(categories []string, posts *[]modules.Post, page int, forumDB *sql.DB) error {
	const limit = 10
	offset := (page - 1) * limit

	categoriesCode := GetCategoryMask(categories)
	categoryInt, err := strconv.ParseInt(categoriesCode, 2, 64)
	if err != nil {
		return fmt.Errorf("error parsing category mask: %v", err)
	}

	query := `SELECT user_id, item_id, title, content, categories, created_at 
              FROM posts 
              WHERE (CAST(categories AS INTEGER) & ?) != 0 
              ORDER BY created_at DESC 
              LIMIT ? OFFSET ?`

	rows, err := forumDB.Query(query, categoryInt, limit, offset)
	if err != nil {
		return fmt.Errorf("error querying posts by category: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post modules.Post
		var categoryMask string

		if err := rows.Scan(&post.Publisher.Username, &post.ID, &post.Title, &post.Text, &categoryMask, &post.CreationTime); err != nil {
			return fmt.Errorf("error scanning post: %v", err)
		}

		post.Categories = GetCategoriesFromMask(categoryMask)
		*posts = append(*posts, post)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

func fetchAllPosts(posts *[]modules.Post, page int, forumDB *sql.DB) error {
	const limit = 10
	offset := (page - 1) * limit

	query := `SELECT user_id, id, title, content, categories, created_at 
              FROM posts 
              ORDER BY created_at DESC 
              LIMIT ? OFFSET ?`

	rows, err := forumDB.Query(query, limit, offset)
	if err != nil {
		return fmt.Errorf("error querying all posts: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post modules.Post
		var categoryMask string

		if err := rows.Scan(&post.Publisher.Id, &post.ID, &post.Title, &post.Text, &categoryMask, &post.CreationTime); err != nil {
			return fmt.Errorf("error scanning post: %v", err)
		}
		err := GetPublicUser(&post.Publisher, forumDB)
		if err != nil {
			fmt.Println(err)
			return err
		}
		post.Categories = GetCategoriesFromMask(categoryMask)

		*posts = append(*posts, post)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
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
