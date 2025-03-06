package internal

import "errors"

var (
	ErrorPageNotFound     = errors.New("page not found")
	ErrorInvalidJSON      = errors.New("request body is not a valid json")
	ErrorInvalidToken     = errors.New("bearer token is invalid")
	ErrorInvalidLogMethod = errors.New("invalid log method. only 'read' and 'write' accepted")
	ErrorInvalidRequest   = errors.New("invalid request")
	ErrorInvalidUserIp    = errors.New("invalid user ip address")
	ErrorInvalidUserId    = errors.New("invalid user id")
	ErrorInvalidAction    = errors.New("action property cannot be empty")
	ErrorInvalidUserAgent = errors.New("user agent property cannot be empty")
	ErrorInvalidContent   = errors.New("content property is invalid. please enter valid JSON object")
)
