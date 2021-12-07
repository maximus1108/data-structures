package main

import (
	"fmt"
)

type Node struct {
	value int // for simplicity for this demo use ints
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Append(value int) {
	newNode := &Node{value: value}
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

func (l *LinkedList) Prepend(value int) {
	newNode := &Node{value: value}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList) Delete(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		return
	}

	node := l.head
	for node != nil && node.next.value != value {
		node = node.next
	}

	if node != nil {
		node.next = node.next.next
	}

}

func (l *LinkedList) DeleteHead() {
	if l.head != nil {
		l.head = l.head.next
	}
}

func (l *LinkedList) Find(value int) *Node {
	node := l.head
	for node != nil && node.value != value {
		node = node.next
	}
	return node
}

func (l *LinkedList) NthNode(i int) *Node {
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

func (l *LinkedList) NthValue(i int) (int, error) {
	if node := l.NthNode(i); node != nil {
		return node.value, nil
	}
	return 0, fmt.Errorf("%d is not a valid index", i)
}

func (l *LinkedList) InsertAfter(node *Node, value int) {
	if node != nil {
		newNode := &Node{value: value}
		newNode.next = node.next
		node.next = newNode
	}
}

func (l *LinkedList) InsertAt(i int, value int) {
	if i < 0 {
		return
	}
	newNode := &Node{value: value}
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

func (l *LinkedList) Reverse() {
	node := l.head
	var nextNode *Node
	var prevNode *Node
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
