package doublylinkedlist

type Node[T comparable] struct {
	Value T
	next  *Node[T]
	prev  *Node[T]
}

type DoublyLinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

func New[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) Prepend(value T) {
	node := &Node[T]{Value: value, next: l.head}
	if l.head != nil {
		l.head.prev = node
	}

	l.head = node

	if l.tail == nil {
		l.tail = node
	}
}

func (l *DoublyLinkedList[T]) Append(value T) {
	node := &Node[T]{Value: value, prev: l.tail}
	if l.tail != nil {
		l.tail.next = node
	}

	l.tail = node

	if l.head == nil {
		l.head = node
	}
}

func (l *DoublyLinkedList[T]) DeleteFirst(value T) *Node[T] {

	node := l.head
	for node != nil && node.Value != value {
		node = node.next
	}

	if node == nil {
		return nil
	}

	if node == l.head {
		l.head = nil
		if node.next != nil {
			l.head = node.next
			l.head.prev = nil
		}
		if node == l.tail {
			l.tail = nil
		}
	} else if node == l.tail {
		l.tail = node.prev
		l.tail.next = nil
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	return node
}

func (l *DoublyLinkedList[T]) DeleteAll(value T) {
	node := l.head
	for node != nil {
		if node.Value != value {
			node = node.next
			continue
		}

		if node == l.head {
			l.head = nil

			if node.next != nil {
				l.head = node.next
				l.head.prev = nil
			}
			if node == l.tail {
				l.tail = nil
			}
		} else if node == l.tail {
			l.tail = node.prev
			l.tail.next = nil
		} else {
			node.next.prev = node.prev
			node.prev.next = node.next
		}
		node = node.next
	}
}

func (l *DoublyLinkedList[T]) DeleteTail() *Node[T] {
	if l.tail == nil {
		return nil
	}

	deletedTail := l.tail
	if l.tail == l.head {
		l.tail = nil
		l.head = nil
		return deletedTail
	}

	l.tail = deletedTail.prev
	l.tail.next = nil
	deletedTail.prev = nil
	return deletedTail
}

func (l *DoublyLinkedList[T]) DeleteHead() *Node[T] {
	if l.head == nil {
		return nil
	}

	deletedHead := l.head
	if l.tail == l.head {
		l.tail = nil
		l.head = nil
		return deletedHead
	}

	l.head = deletedHead.next
	l.head.prev = nil
	deletedHead.next = nil
	return deletedHead
}

func (l *DoublyLinkedList[T]) Head() *Node[T] {
	return l.head
}

func (l *DoublyLinkedList[T]) Tail() *Node[T] {
	return l.tail
}

func (l *DoublyLinkedList[T]) Find(value T) *Node[T] {
	node := l.head
	for node != nil && node.Value != value {
		node = node.next
	}
	return node
}

func (l *DoublyLinkedList[T]) NthNode(i uint) *Node[T] {
	node := l.head
	for i > 0 && node != nil {
		node = node.next
		i--
	}
	return node
}

func (l *DoublyLinkedList[T]) NthValue(i uint) (T, bool) {
	if node := l.NthNode(i); node != nil {
		return node.Value, true
	}
	var zero T
	return zero, false
}

func (l *DoublyLinkedList[T]) InsertAfter(node *Node[T], value T) {
	if node == nil || l.head == nil {
		return
	}

	newNode := &Node[T]{Value: value}
	newNode.next = node.next
	newNode.prev = node
	if node == l.tail {
		l.tail = newNode
	}
	if node.next != nil {
		node.next.prev = newNode
	}
	node.next = newNode
}

func (l *DoublyLinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == nil || l.head == nil {
		return
	}

	newNode := &Node[T]{Value: value}
	newNode.next = node
	newNode.prev = node.prev
	if node == l.head {
		l.head = newNode
	}

	if node.prev != nil {
		node.prev.next = newNode
	}

	node.prev = newNode

}

func (l *DoublyLinkedList[T]) InsertAt(i uint, value T) {

	newNode := &Node[T]{Value: value}
	if i == 0 {
		if l.head == nil {
			l.head = newNode
			l.tail = newNode
			return
		}
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
		return
	}

	if l.head == nil {
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
	if node == l.tail {
		l.tail = newNode
	}
}

func (l *DoublyLinkedList[T]) Reverse() {
	node := l.head
	var nextNode *Node[T]
	var prevNode *Node[T]
	for node != nil {
		// get the next node
		nextNode = node.next
		// update the current nodes pointers
		node.next = prevNode
		node.prev = nextNode
		// move to the next node..
		prevNode = node
		node = nextNode
	}
	l.tail = l.head
	l.head = prevNode
}
