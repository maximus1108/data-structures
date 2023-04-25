package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New[string]()

	for i := 0; i < 10; i++ {
		fmt.Println("Enqueueing", i)
		q.Enqueue(fmt.Sprintf("value %d", i))
	}

	for i := 0; i < 10; i++ {
		expected := fmt.Sprintf("value %d", i)
		fmt.Println("expecting", i)
		actual := q.Dequeue()
		fmt.Println("got", actual)

		if actual != expected {
			t.Errorf("expected '%s', got '%s'", expected, actual)
		}
	}

	if !q.IsEmpty() {
		t.Error("expected queue to be empty")
	}

}
