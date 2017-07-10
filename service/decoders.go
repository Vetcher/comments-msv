package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeRequestOnlyWithID(_ context.Context, r *http.Request) (interface{}, error) {
	var request RequestOnlyWithId
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return request, nil
}

func DecodeRequestPostComment(_ context.Context, r *http.Request) (interface{}, error) {
	var request RequestPostComment
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("parsing error: %v", err)
	}
	return request, nil
}
