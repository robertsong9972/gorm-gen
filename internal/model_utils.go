package codegen

import (
	"errors"
	"strings"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrEmptyQuery   = errors.New("empty query")
)

type Cond struct {
	Cond string
	Op   string
}

func IgnoreMessageName(messageName string) bool {
	// black list these world
	if messageName == "Constant" || strings.HasSuffix(messageName, "Response") {
		return true
	}
	return false
}
