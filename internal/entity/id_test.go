package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID_TypeAndPureID(t *testing.T) {
	id := NewID("testType", "testID")

	assert.Equal(t, "testType", id.Type())
	assert.Equal(t, "testID", id.PureID())
}
