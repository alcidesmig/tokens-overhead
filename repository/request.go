package repository

type RequestInterface interface {
	Request(token, address string) error
}
