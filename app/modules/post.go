package modules

import (
	"net/http"
	"time"

	"forum/app/config"
	"forum/app/modules/errors"
)

type PostContent struct {
	Title      string   `json:"title"`
	Text       string   `json:"text"`
	Image      string   `json:"image,omitempty"`
	Categories []string `json:"categories"`
}
type Post struct {
	Id           int         `json:"id"`
	Content      PostContent `json:"content"`
	Likes        int         `json:"likes"`
	Dislikes     int         `json:"dislikes"`
	Reaction     string      `json:"reaction"`
	CreationTime time.Time   `json:"creationTime"`
	Publisher    User        `json:"publisher"`
}

func (pc *PostContent) ValidatePostContent() (err *errors.HttpError) {
	if len(pc.Title) == 0 || len([]rune(pc.Title)) > 50 {
		return errors.NewError(http.StatusBadRequest, errors.CodeInvalidRequestFormat, "title too long", "title can't be empty or more than 50 character")
	}
	if len(pc.Text) == 0 || len([]rune(pc.Text)) > 5000 {
		return errors.NewError(http.StatusBadRequest, errors.CodeInvalidRequestFormat, "content too long", "content can't be empty or more than 5000 character")
	}
	if len(pc.Categories) > config.MaxCategoriesSize {
		return errors.NewError(http.StatusBadRequest, errors.CodeInvalidRequestFormat, "too many categories", "can't select more than 4 categories")
	}
	return nil
}
