package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New[int]()

	for i := 0; i < 10; i++ {
		q.Enqueue(i)

	}

	for i := 0; i < 10; i++ {
		if q.Peek() != i {
			t.Errorf("expected %d, got %d", i, q.Peek())
		}
		if q.Dequeue() != i {
			t.Errorf("expected %d, got %d", i, q.Dequeue())
		}
	}

	if !q.IsEmpty() {
		t.Error("expected queue to be empty")
	}

}
