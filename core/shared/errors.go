package shared

import "errors"

var (
	ErrorUnsupportedType = errors.New("unsupported type")
	ErrorEmptyResponse   = errors.New("server response was empty")
)
