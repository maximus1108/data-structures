package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	t.Run("should have the expected elements in the queue when enqueuing and dequeuing", func(t *testing.T) {
		q := New[int]()

		for i := 0; i < 10; i++ {
			q.Enqueue(i)
		}

		for i := 0; i < 10; i++ {
			assert.Equal(t, i, q.Peek())
			v, ok := q.Dequeue()
			assert.True(t, ok)

			assert.Equal(t, i, v)
		}

		assert.True(t, q.IsEmpty())
		v, ok := q.Dequeue()
		assert.False(t, ok)
		assert.Equal(t, 0, v)
	})
}
