package service

type RequestGetCommentByID struct {
	ID uint `json:"id"`
}

type RequestGetCommentByAuthorID struct {
	ID uint `json:"id"`
}

type RequestDeleteCommentByID struct {
	ID uint `json:"id"`
}

type RequestPostComment struct {
	Text     string `json:"text"`
	AuthorID uint   `json:"author_id"`
}
