package linkedlist

import (
	"testing"
)

func assertLinkedListValueSequenceMatchesSlice(t *testing.T, expected []int, actual LinkedList[int]) {
	node := actual.head
	for i, expectedValue := range expected {
		if node == nil {
			t.Errorf("not enough nodes in linkedlist, expected %d", len(expected))
			return
		}
		if node.Value != expectedValue {
			t.Errorf("expected %d at index %d, but %d", expectedValue, i, node.Value)
		}
		node = node.next
	}
	if node != nil {
		t.Errorf("too many nodes in linkedlist, expected %d", len(expected))
	}
}

func assertNotNilOrIncorrectValue(t *testing.T, expected int, actual *Node[int]) {
	if actual == nil || actual.Value != expected {
		t.Errorf("expected %d but got %v", expected, actual)
	}
}
func TestAppendShouldAddValuesToTheEndOfTheLinkedList(t *testing.T) {
	list := LinkedList[int]{}
	Values := [5]int{1, 5, 3, 17, 66}

	for _, Value := range Values {
		list.Append(Value)
	}

	assertLinkedListValueSequenceMatchesSlice(t, Values[:], list)
}

func TestPrependShouldAddValuesToTheBeginningOfTheLinkedList(t *testing.T) {
	list := LinkedList[int]{}
	Values := [5]int{11, 2, 8, 77, 4}

	for _, Value := range Values {
		list.Prepend(Value)
	}

	assertLinkedListValueSequenceMatchesSlice(t, []int{4, 77, 8, 2, 11}, list)

}

func TestDeleteShouldRemoveFirstMatchingValueFromTheLinkedList(t *testing.T) {
	list := LinkedList[int]{}
	Values := [6]int{11, 2, 5, 8, 77, 5}

	for _, Value := range Values {
		list.Append(Value)
	}
	list.Delete(5)
	list.Delete(11)
	list.Delete(77)
	list.Delete(5)

	assertLinkedListValueSequenceMatchesSlice(t, []int{2, 8}, list)

}

func TestDeleteFirstShouldRemoveFirstFromElementTheLinkedList(t *testing.T) {
	list := LinkedList[int]{}
	Values := [6]int{11, 2, 5, 8, 77, 5}

	for _, Value := range Values {
		list.Append(Value)
	}

	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, Values[1:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, Values[2:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, Values[3:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, Values[4:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, Values[5:], list)
	list.DeleteHead()
	assertLinkedListValueSequenceMatchesSlice(t, Values[6:], list)
}

func TestInsertAfterShouldInsertValueAfterSpecifiedNode(t *testing.T) {
	list := LinkedList[int]{}
	nodeA := &Node[int]{Value: 11}
	list.head = nodeA
	nodeB := &Node[int]{Value: 2}
	nodeA.next = nodeB
	nodeC := &Node[int]{Value: 5}
	nodeB.next = nodeC
	nodeD := &Node[int]{Value: 6}
	nodeC.next = nodeD

	list.InsertAfter(nodeB, 7)
	list.InsertAfter(nodeB, 3)
	list.InsertAfter(nodeD, 99)

	assertLinkedListValueSequenceMatchesSlice(t, []int{11, 2, 3, 7, 5, 6, 99}, list)

}
func TestInsertAtShouldInsertAtValueSpecifiedIndex(t *testing.T) {
	list := LinkedList[int]{}
	Values := [6]int{11, 2, 5, 8, 77, 5}

	for _, Value := range Values {
		list.Append(Value)
	}

	list.InsertAt(0, 81)
	list.InsertAt(2, 99)
	list.InsertAt(88, 14) // out of bounds index should do nothing
	list.InsertAt(7, 64)

	assertLinkedListValueSequenceMatchesSlice(t, []int{81, 11, 99, 2, 5, 8, 77, 64, 5}, list)

}

func TestReverseShouldReverseTheSequenceOfValues(t *testing.T) {
	list := LinkedList[int]{}
	Values := [6]int{11, 2, 5, 8, 77, 5}

	for _, Value := range Values {
		list.Append(Value)
	}

	list.Reverse()

	assertLinkedListValueSequenceMatchesSlice(t, []int{5, 77, 8, 5, 2, 11}, list)

}

func TestSearchShouldFindNodeOrNil(t *testing.T) {
	list := LinkedList[int]{}
	Values := [6]int{11, 2, 5, 8, 77, 5}

	for _, Value := range Values {
		list.Append(Value)
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
	list := LinkedList[int]{}
	Values := [6]int{11, 2, 5, 8, 77, 5}
	for _, Value := range Values {
		list.Append(Value)
	}

	for i, Value := range Values {
		assertNotNilOrIncorrectValue(t, Value, list.NthNode(i))
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
	list := LinkedList[int]{}
	Values := [6]int{11, 2, 5, 8, 77, 5}
	for _, Value := range Values {
		list.Append(Value)
	}

	for i, Value := range Values {
		result, err := list.NthValue(i)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
			continue
		}
		if result != Value {
			t.Errorf("result %d at indexs %d, but got %d", Value, i, result)
		}
	}

	Value, err := list.NthValue(77)
	if Value != 0 || err == nil {
		t.Errorf("expected 0 and an error but got %d and '%s'", Value, err)
	}
	Value, err = list.NthValue(-1)
	if Value != 0 || err == nil {
		t.Errorf("expected 0 and an error but got %d and '%s'", Value, err)
	}
}
