package internal

import "errors"

var (
	ErrorPageNotFound     = errors.New("page not found")
	ErrorInvalidJSON      = errors.New("request body is not a valid json")
	ErrorInvalidToken     = errors.New("bearer token is invalid")
	ErrorInvalidLogMethod = errors.New("invalid log method. only 'read' and 'write' accepted")
)
