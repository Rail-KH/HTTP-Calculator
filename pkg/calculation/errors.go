package calculation

import (
	"errors"
)

var (
	UnprocessableEntity = errors.New("Unprocessable Entity")
	InternalServerError = errors.New("Internal Server Error")
)
