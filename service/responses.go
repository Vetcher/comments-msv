package service

type JsonResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error,omitempty"`
}
