package linkedlist

type Node[T comparable] struct {
	Value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
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

func (l *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}
	newNode.next = l.head
	l.head = newNode
}

func (l *LinkedList[T]) DeleteFirst(value T) *Node[T] {

	if l.head != nil && l.head.Value == value {
		oldHead := l.head
		l.head = l.head.next
		return oldHead
	}

	nodeBefore := l.head
	for nodeBefore != nil && nodeBefore.next != nil {
		if nodeBefore.next.Value != value {
			deleted := nodeBefore.next
			nodeBefore.next = nodeBefore.next.next
			return deleted
		}
		nodeBefore = nodeBefore.next
	}
	return nil
}

func (l *LinkedList[T]) DeleteAll(value T) {

	if l.head != nil && l.head.Value == value {
		l.head = l.head.next
	}

	nodeBefore := l.head
	for nodeBefore != nil && nodeBefore.next != nil {
		if nodeBefore.next.Value == value {
			nodeBefore.next = nodeBefore.next.next
		}
		nodeBefore = nodeBefore.next
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

func (l *LinkedList[T]) Find(value T) *Node[T] {
	node := l.head
	for node != nil && node.Value != value {
		node = node.next
	}
	return node
}

func (l *LinkedList[T]) Head() *Node[T] {
	return l.head
}

func (l *LinkedList[T]) NthNode(i uint) *Node[T] {
	node := l.head
	for i > 0 && node != nil {
		node = node.next
		i--
	}
	return node
}

func (l *LinkedList[T]) NthValue(i uint) (T, bool) {
	if node := l.NthNode(i); node != nil {
		return node.Value, true
	}
	var zero T
	return zero, false
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if node != nil {
		newNode := &Node[T]{Value: value}
		newNode.next = node.next
		node.next = newNode
	}
}

func (l *LinkedList[T]) InsertAt(i uint, value T) {

	newNode := &Node[T]{Value: value}
	if i == 0 {
		if l.head == nil {
			l.head = newNode
			return
		}
		newNode.next = l.head
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
