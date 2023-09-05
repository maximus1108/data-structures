package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	t.Run("should have the expected elements in the stack when pushing and popping", func(t *testing.T) {
		s := New[int]()

		assert.True(t, s.IsEmpty())

		for i := 0; i < 10; i++ {
			s.Push(i)
			assert.Equal(t, i, s.Peek())
		}

		for i := 9; i >= 0; i-- {
			assert.Equal(t, i, s.Peek())
			assert.Equal(t, i, s.Pop())
		}

		assert.True(t, s.IsEmpty())
	})
}
