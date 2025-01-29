package modules

import (
	"time"
)

type Post struct {
	Title        string `json:"title"`
	Image        string `json:"image,omitempty"`
	Text         string `json:"text"`
	ID           string
	Categories   string  `json:"categories"`
	CreationTime time.Time `json:"creationTime"`
	Publisher    User      `json:"publisher"`
}
