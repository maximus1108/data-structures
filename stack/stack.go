package stack

import "github.com/maximus1108/data-structures/linkedlist"

func New[T comparable]() *Stack[T] {
	return &Stack[T]{
		elements: linkedlist.New[T](),
	}
}

type Stack[T comparable] struct {
	elements *linkedlist.LinkedList[T]
}

func (s *Stack[T]) Push(value T) {
	s.elements.Prepend(value)
}

func (s *Stack[T]) Pop() T {
	node := s.elements.DeleteHead()
	if node == nil {
		var zero T
		return zero
	}
	return node.Value
}

func (s *Stack[T]) Peek() T {
	head := s.elements.Head()
	if head == nil {
		var zero T
		return zero
	}
	return head.Value
}

func (s *Stack[T]) IsEmpty() bool {
	return s.elements.Head() == nil
}
