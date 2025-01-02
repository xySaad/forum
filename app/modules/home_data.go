package modules

type Post struct {
	Title      string   `json:"title"`
	Image      string   `json:"image"`
	Text       string   `json:"text"`
	Categories []string `json:"categories"`
}
