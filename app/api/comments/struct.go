package comments

type Comment struct {
	UserID    string `json:"user_id"`
	PostID    string `json:"post_id"`
	ItemID    string `json:"item_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}
