package comments

import "forum/app/modules"

type Comment struct {
	UserID    string `json:"user_id"`
	Publisher modules.User `json:"publisher"`
	PostID    string `json:"post_id"`
	ItemID    string `json:"item_id"`
	Content   string `json:"content"`
	Likes     int    `json:"likes"`
	Dislikes  int    `json:"dislike"`
	CreatedAt string `json:"created_at"`
}
