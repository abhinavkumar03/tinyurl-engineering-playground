package errors

import "errors"

var (
	ErrURLNotFound = errors.New("url not found")

	ErrInvalidURL = errors.New("invalid url")

	ErrInternalServer = errors.New("internal server error")

	ErrCacheMiss = errors.New("cache miss")
)
