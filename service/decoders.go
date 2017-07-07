package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestOnlyWithId struct {
	ID uint `json:"id"`
}

func DecodeRequestOnlyWithId(_ context.Context, r *http.Request) (interface{}, error) {
	var request RequestOnlyWithId
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return request, nil
}

type RequestPostComment struct {
	Text     string `json:"text"`
	AuthorID uint   `json:"author_id"`
}

func DecodeRequestPostComment(_ context.Context, r *http.Request) (interface{}, error) {
	var request RequestPostComment
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return request, nil
}
