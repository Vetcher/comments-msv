package service

type RequestOnlyWithID struct {
	ID uint `json:"id"`
}

type RequestPostComment struct {
	Text     string `json:"text"`
	AuthorID uint   `json:"author_id"`
}