package linkedlist

import (
	"fmt"
)

type Node[T comparable] struct {
	Value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Append(Value T) {
	newNode := &Node[T]{Value: Value}
	if l.head == nil {
		l.head = newNode
		return
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = newNode
}

func (l *LinkedList[T]) Prepend(Value T) {
	newNode := &Node[T]{Value: Value}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList[T]) Delete(Value T) {
	if l.head == nil {
		return
	}

	if l.head.Value == Value {
		l.head = l.head.next
		return
	}

	node := l.head
	for node != nil && node.next.Value != Value {
		node = node.next
	}

	if node != nil {
		node.next = node.next.next
	}

}

func (l *LinkedList[T]) DeleteHead() *Node[T] {
	if l.head != nil {
		node := l.head
		l.head = l.head.next
		return node
	}
	return nil
}

func (l *LinkedList[T]) Find(Value T) *Node[T] {
	node := l.head
	for node != nil && node.Value != Value {
		node = node.next
	}
	return node
}

func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

func (l *LinkedList[T]) NthNode(i int) *Node[T] {
	if i < 0 {
		return nil
	}
	node := l.head
	for i > 0 && node != nil {
		node = node.next
		i--
	}
	return node
}

func (l *LinkedList[T]) NthValue(i int) (T, error) {
	if node := l.NthNode(i); node != nil {
		return node.Value, nil
	}
	var zero T
	return zero, fmt.Errorf("%d is not a valid index", i)
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], Value T) {
	if node != nil {
		newNode := &Node[T]{Value: Value}
		newNode.next = node.next
		node.next = newNode
	}
}

func (l *LinkedList[T]) InsertAt(i int, Value T) {
	if i < 0 {
		return
	}
	newNode := &Node[T]{Value: Value}
	if i == 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}
	i--
	node := l.head
	for i > 0 {
		node = node.next
		if node == nil {
			return
		}
		i--
	}
	nextNode := node.next
	node.next = newNode
	newNode.next = nextNode
}

func (l *LinkedList[T]) Reverse() {
	node := l.head
	var nextNode *Node[T]
	var prevNode *Node[T]
	for node != nil {
		// get the next node
		nextNode = node.next
		// update the current nodes pointer to the previous
		node.next = prevNode
		// move to the next node..
		prevNode = node
		node = nextNode
	}
	l.head = prevNode
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}
