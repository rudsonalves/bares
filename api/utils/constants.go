package utils

// HttpMethod define um tipo para representar métodos HTTP como enum
type HttpMethod string

// Constantes que representam os métodos HTTP.
const (
	HttpMethodGet    HttpMethod = "GET"
	HttpMethodPost   HttpMethod = "POST"
	HttpMethodPut    HttpMethod = "PUT"
	HttpMethodDelete HttpMethod = "DELETE"
)
