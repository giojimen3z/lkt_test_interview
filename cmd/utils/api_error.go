package utils

type APIError struct {
	Error   string `json:"error"`
	Details string `json:"details"`
}
