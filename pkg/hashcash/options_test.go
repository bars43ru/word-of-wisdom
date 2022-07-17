package hashcash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WithComplexity(t *testing.T) {
	s := New(WithComplexity(5))
	assert.Equal(t, uint8(5), s.complexity)
}
