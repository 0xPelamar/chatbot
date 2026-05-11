package entity

import (
	"fmt"
	"strings"
)

type ID string

func NewID[T any](type_ string, id T) ID {
	return ID(fmt.Sprintf("%s:%v", type_, id))
}

func (id ID) String() string {
	return string(id)
}

func (id ID) Type() string {
	return strings.Split(id.String(), ":")[0]
}

func (id ID) PureID() string {
	return strings.Split(id.String(), ":")[1]
}
