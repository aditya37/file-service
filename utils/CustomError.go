package utils

import (
	"fmt"
)

type CustomError struct {
	InternalError string `json:"internal_error,omitempty"`
	Description   string `json:"description,omitempty"`
	Code          int    `json:"code"`
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("%s", c.InternalError)
}
