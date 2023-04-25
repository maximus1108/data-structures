package queue

import "github.com/maximus1108/data-structures/linkedlist"

func New[T comparable]() *Queue[T] {
	return &Queue[T]{
		elements: &linkedlist.LinkedList[T]{}, // TODO: use doubly linked list to optimize Enqueue
	}
}

type Queue[T comparable] struct {
	elements *linkedlist.LinkedList[T]
}

func (q *Queue[T]) Enqueue(value T) {
	q.elements.Append(value)
}

func (q *Queue[T]) Dequeue() T {
	head := q.elements.DeleteHead()
	if head == nil {
		var zero T
		return zero
	}
	return head.Value
}

func (q *Queue[T]) Peek() T {
	return q.elements.Head().Value
}

func (q *Queue[T]) IsEmpty() bool {
	return q.elements.Head() == nil
}
