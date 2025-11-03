package entity

import (
	"fmt"
	"strings"
)

type ID string

func NewID[T any](entityType string, id T) ID {
	return ID(fmt.Sprintf("%s:%v", entityType, id))
}

func (id ID) String() string {
	return string(id)
}
func (id ID) Type() string {
	return strings.Split(id.String(), ":")[0]
}
func (id ID) ID() string {
	return strings.Split(id.String(), ":")[1]
}
