package common

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type Response[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}
