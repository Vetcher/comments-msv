package service

type RequestOnlyWithId struct {
	ID uint `json:"id"`
}

type RequestPostComment struct {
	Text     string `json:"text"`
	AuthorID uint   `json:"author_id"`
}