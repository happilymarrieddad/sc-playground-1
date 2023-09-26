package types

import (
	"fmt"
	"strings"
)

const (
	notFoundErr     string = "Not Found:"
	unauthorizedErr string = "Unauthorized:"
)

func IsNotFoundError(err error) bool {
	return strings.Contains(err.Error(), notFoundErr)
}

func NewNotFoundError(msg string) error {
	return fmt.Errorf("%s %s", notFoundErr, msg)
}

func IsUnauthorizedError(err error) bool {
	return strings.Contains(err.Error(), unauthorizedErr)
}

func NewUnauthorizedError(msg string) error {
	return fmt.Errorf("%s %s", unauthorizedErr, msg)
}
