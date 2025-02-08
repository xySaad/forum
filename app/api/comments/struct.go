package comments

import "forum/app/modules"

type Comment struct {
	UserID       string       `json:"user_id"`
	Publisher    modules.User `json:"publisher"`
	PostID       string       `json:"post_id"`
	ItemID       int       `json:"item_id"`
	Content      string       `json:"content"`
	Likes        int          `json:"likes"`
	Dislikes     int          `json:"dislikes"`
	Reaction     int          `json:"reaction"`
	CreationTime string       `json:"creationTime"`
}
