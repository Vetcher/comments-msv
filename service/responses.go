package service

type Response struct {
	Data interface{} `json:"data"`
	Err  error       `json:"error,omitempty"`
}
