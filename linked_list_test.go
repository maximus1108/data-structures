package main

import (
	"testing"
)

func assertLinkedListValueSequenceMatchesSlice(t *testing.T, expected []int, actual LinkedList) {
	node := actual.head
	for i, expectedValue := range expected {
		if node == nil {
			t.Errorf("not enough nodes in linkedlist, expected %d", len(expected))
			return
		}
		if node.value != expectedValue {
			t.Errorf("expected %d at index %d, but %d", expectedValue, i, node.value)
		}
		node = node.next
	}
	if node != nil {
		t.Errorf("too many nodes in linkedlist, expected %d", len(expected))
	}
}

func assertNotNilOrIncorrectValue(t *testing.T, expected int, actual *Node) {
	if actual == nil || actual.value != expected {
		t.Errorf("expected %d but got %v", expected, actual)
	}
}
func TestAppendShouldAddValuesToTheEndOfTheLinkedList(t *testing.T) {
	list := LinkedList{}
	values := [5]int{1, 5, 3, 17, 66}

	for _, value := range values {
		list.Append(value)
	}

	assertLinkedListValueSequenceMatchesSlice(t, values[:], list)
}

func TestPrependShouldAddValuesToTheBeginningOfTheLinkedList(t *testing.T) {
	list := LinkedList{}
	values := [5]int{11, 2, 8, 77, 4}

	for _, value := range values {
		list.Prepend(value)
	}

	assertLinkedListValueSequenceMatchesSlice(t, []int{4, 77, 8, 2, 11}, list)

}

func TestDeleteShouldRemoveFirstMatchingValueFromTheLinkedList(t *testing.T) {
	list := LinkedList{}
	values := [6]int{11, 2, 5, 8, 77, 5}

	for _, value := range values {
		list.Append(value)
	}
	list.Delete(5)
	list.Delete(11)
	list.Delete(77)
	list.Delete(5)

	assertLinkedListValueSequenceMatchesSlice(t, []int{2, 8}, list)

}

func TestDeleteFirstShouldRemoveFirstFromElementTheLinkedList(t *testing.T) {
	list := LinkedList{}
	values := [6]int{11, 2, 5, 8, 77, 5}

	for _, value := range values {
		list.Append(value)
	}

	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, values[1:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, values[2:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, values[3:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, values[4:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, values[5:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, values[6:], list)
}

func TestInsertAfterShouldInsertValueAfterSpecifiedNode(t *testing.T) {
	list := LinkedList{}
	nodeA := &Node{value: 11}
	list.head = nodeA
	nodeB := &Node{value: 2}
	nodeA.next = nodeB
	nodeC := &Node{value: 5}
	nodeB.next = nodeC
	nodeD := &Node{value: 6}
	nodeC.next = nodeD

	list.InsertAfter(nodeB, 7)
	list.InsertAfter(nodeB, 3)
	list.InsertAfter(nodeD, 99)

	assertLinkedListValueSequenceMatchesSlice(t, []int{11, 2, 3, 7, 5, 6, 99}, list)

}
func TestInsertAtShouldInsertAtValueSpecifiedIndex(t *testing.T) {
	list := LinkedList{}
	values := [6]int{11, 2, 5, 8, 77, 5}

	for _, value := range values {
		list.Append(value)
	}

	list.InsertAt(0, 81)
	list.InsertAt(2, 99)
	list.InsertAt(88, 14) // out of bounds index should do nothing
	list.InsertAt(7, 64)

	assertLinkedListValueSequenceMatchesSlice(t, []int{81, 11, 99, 2, 5, 8, 77, 64, 5}, list)

}

func TestReverseShouldReverseTheSequenceOfValues(t *testing.T) {
	list := LinkedList{}
	values := [6]int{11, 2, 5, 8, 77, 5}

	for _, value := range values {
		list.Append(value)
	}

	list.Reverse()

	assertLinkedListValueSequenceMatchesSlice(t, []int{5, 77, 8, 5, 2, 11}, list)

}

func TestSearchShouldFindNodeOrNil(t *testing.T) {
	list := LinkedList{}
	values := [6]int{11, 2, 5, 8, 77, 5}

	for _, value := range values {
		list.Append(value)
	}

	validSearchValues := [3]int{11, 77, 5}
	for _, searchValue := range validSearchValues {
		assertNotNilOrIncorrectValue(t, searchValue, list.Find(searchValue))
	}

	node := list.Find(876)
	if node != nil {
		t.Errorf("expected nil but got %v", node)
	}

}

func TestNthNodeShouldReturnNodeOrNil(t *testing.T) {
	list := LinkedList{}
	values := [6]int{11, 2, 5, 8, 77, 5}
	for _, value := range values {
		list.Append(value)
	}

	for i, value := range values {
		assertNotNilOrIncorrectValue(t, value, list.NthNode(i))
	}

	node := list.NthNode(77)
	if node != nil {
		t.Errorf("expected nil but got %v", node)
	}
	node = list.NthNode(-1)
	if node != nil {
		t.Errorf("expected nil but got %v", node)
	}
}
func TestNthValueShouldReturnValueOrError(t *testing.T) {
	list := LinkedList{}
	values := [6]int{11, 2, 5, 8, 77, 5}
	for _, value := range values {
		list.Append(value)
	}

	for i, value := range values {
		result, err := list.NthValue(i)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			continue
		}
		if result != value {
			t.Errorf("result %d at indexs %d, but got %d", value, i, result)
		}
	}

	value, err := list.NthValue(77)
	if value != 0 || err == nil {
		t.Errorf("expected 0 and an error but got %d and '%s'", value, err)
	}
	value, err = list.NthValue(-1)
	if value != 0 || err == nil {
		t.Errorf("expected 0 and an error but got %d and '%s'", value, err)
	}
}
