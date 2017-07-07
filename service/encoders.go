package service

import (
	"context"
	"encoding/json"
	"net/http"
)

func ServeJSON(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
