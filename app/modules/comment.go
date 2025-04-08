package modules

type Comment struct {
	Publisher    User   `json:"publisher"`
	PostId       string `json:"postId"`
	Id           string `json:"id"`
	Content      string `json:"content"`
	Likes        int    `json:"likes"`
	Dislikes     int    `json:"dislikes"`
	Reaction     string `json:"reaction"`
	CreationTime string `json:"creationTime"`
}
