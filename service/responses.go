package service

type Response struct {
	Data interface{} `json:"data"`
	Err  string      `json:"error,omitempty"`
}
