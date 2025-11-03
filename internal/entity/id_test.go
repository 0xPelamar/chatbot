package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestID_TypeAndValue(t *testing.T) {
	assert.Equal(t, "type", NewID("type", "value").Type())
	assert.Equal(t, "value", NewID("type", "value").ID())
}
