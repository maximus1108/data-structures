package queue

import (
	"github.com/maximus1108/data-structures/doublylinkedlist"
)

func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		elements: doublylinkedlist.New[T](),
	}
}

type Queue[T comparable] struct {
	elements *doublylinkedlist.DoublyLinkedList[T]
}

func (q *Queue[T]) Enqueue(value T) {
	q.elements.Append(value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	head := q.elements.DeleteHead()
	if head == nil {
		var zero T
		return zero, false
	}
	return head.Value, true
}

func (q *Queue[T]) Peek() T {
	head := q.elements.Head()
	if head == nil {
		var zero T
		return zero
	}
	return head.Value
}

func (q *Queue[T]) IsEmpty() bool {
	return q.elements.Head() == nil
}
