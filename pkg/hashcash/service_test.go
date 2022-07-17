package hashcash

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Solve(t *testing.T) {
	ctx := context.Background()
	s := New(WithComplexity(0))
	r, err := s.Solve(ctx, []byte("test data"))
	assert.NoError(t, err)
	assert.Equal(t, []byte{}, r)

	s = New(WithComplexity(5))
	r, err = s.Solve(ctx, []byte("test data"))
	assert.NoError(t, err)
	assert.Equal(t, []byte("MTAwMjEwMA=="), r)
}

func Test_SolveCancelContext(t *testing.T) {
	s := New(WithComplexity(0))
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	r, err := s.Solve(ctx, []byte("test data"))
	assert.ErrorIs(t, ctx.Err(), err)
	assert.Equal(t, []byte{}, r)
}

func Test_Verify(t *testing.T) {
	s := New(WithComplexity(0))
	assert.True(t, s.Verify([]byte("test data"), []byte("test data")))

	s = New(WithComplexity(5))
	assert.True(t, s.Verify([]byte("test data"), []byte("MTAwMjEwMA==")))

	s = New(WithComplexity(8))
	assert.False(t, s.Verify([]byte("test data"), []byte("test data")))
}
