package comments

import "forum/app/modules"

type Comment struct {
	UserID       string       `json:"user_id"`
	Publisher    modules.User `json:"publisher"`
	PostID       string       `json:"post_id"`
	ItemID       string       `json:"item_id"`
	Content      string       `json:"content"`
	Reaction     int          `json:"reaction"`
	Likes        int          `json:"likes"`
	Dislikes     int          `json:"dislikes"`
	CreationTime string       `json:"creationTime"`
}
