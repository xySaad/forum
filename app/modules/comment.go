package modules

import "forum/app/modules/snowflake"

type Comment struct {
	Publisher    User                  `json:"publisher"`
	PostId       string                `json:"postId"`
	Id           snowflake.SnowflakeID `json:"id"`
	Content      string                `json:"content"`
	Likes        int                   `json:"likes"`
	Dislikes     int                   `json:"dislikes"`
	Reaction     string                `json:"reaction"`
	CreationTime string                `json:"creationTime"`
}
