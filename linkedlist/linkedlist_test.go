package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList_Append(t *testing.T) {
	l := New[int]()

	values := []int{1, 2, 3, 4}
	for _, n := range values {
		l.Append(n)
	}

	assert.Equal(t, 1, l.head.Value)

	for i, n := range values {
		v, ok := l.NthValue(uint(i))
		assert.True(t, ok)
		assert.Equal(t, n, v)
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	l := New[int]()

	values := []int{1, 2, 3, 4}
	for _, n := range values {
		l.Prepend(n)
	}

	assert.Equal(t, 4, l.head.Value)

	for i := range values {
		n, ok := l.NthValue(uint(i))
		assert.True(t, ok)
		assert.Equal(t, n, values[len(values)-1-i])
	}
}

func TestLinkedList_DeleteFirst(t *testing.T) {

	t.Run("should delete first encountered instance of a specified value", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 2, 4, 2} {
			l.Append(n)
		}

		n := l.DeleteFirst(2)
		assert.Equal(t, 2, n.Value)

		assert.Equal(t, 1, l.head.Value)

		for i, n := range []int{1, 3, 2, 4, 2} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should return nil and not delete any elements if value is not found", func(t *testing.T) {
		l := New[int]()
		values := []int{1, 2, 3, 2, 4, 2}
		for _, n := range values {
			l.Append(n)
		}

		n := l.DeleteFirst(7)
		assert.Nil(t, n)

		assert.Equal(t, 1, l.head.Value)

		for i, n := range values {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should result in an empty list if it only has one element and it matches specified value", func(t *testing.T) {
		l := New[int]()
		values := []int{2}
		for _, n := range values {
			l.Append(n)
		}

		n := l.DeleteFirst(2)
		assert.Equal(t, 2, n.Value)

		assert.Nil(t, l.head)
	})

	t.Run("should update list appropriately if the matching element is the head", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{2, 1, 3, 4} {
			l.Append(n)
		}

		n := l.DeleteFirst(2)
		assert.Equal(t, 2, n.Value)

		assert.Equal(t, 1, l.head.Value)

		for i, n := range []int{1, 3, 4} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

}

func TestLinkedList_DeleteAll(t *testing.T) {

	t.Run("should delete all encountered instances of a specified value", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{2, 1, 2, 3, 2, 4, 2} {
			l.Append(n)
		}

		l.DeleteAll(2)

		assert.Equal(t, 1, l.head.Value)

		for i, n := range []int{1, 3, 4} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should result in an empty list if it only has one element and it matches specified value", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{2} {
			l.Append(n)
		}

		l.DeleteAll(2)

		assert.Nil(t, l.head)

	})

	t.Run("should not delete any elements if value is not found", func(t *testing.T) {
		l := New[int]()
		values := []int{1, 2, 3, 2, 4, 2}
		for _, n := range values {
			l.Append(n)
		}

		l.DeleteAll(7)

		for i, n := range values {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

}

func TestDoublyLinkedList_DeleteHead(t *testing.T) {

	t.Run("should delete head from list", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 4} {
			l.Append(n)
		}

		l.DeleteHead()
		assert.Equal(t, 2, l.head.Value)

		for i, n := range []int{2, 3, 4} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should result in an empty list if it only has one element", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{2} {
			l.Append(n)
		}

		l.DeleteHead()

		assert.Nil(t, l.head)
	})

}
func TestLinkedList_InsertAfter(t *testing.T) {

	t.Run("should insert value after specified node", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 5} {
			l.Append(n)
		}

		n := l.NthNode(2)

		l.InsertAfter(n, 7)

		for i, n := range []int{1, 2, 3, 7, 5} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("new node should become second node if specified node is head", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 5} {
			l.Append(n)
		}

		n := l.Head()

		l.InsertAfter(n, 7)

		for i, n := range []int{1, 7, 2, 3, 5} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should be noop if node is nil", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 5} {
			l.Append(n)
		}

		l.InsertAfter(nil, 7)

		for i, n := range []int{1, 2, 3, 5} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should be noop if specified node is not in list", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 5} {
			l.Append(n)
		}

		n := &Node[int]{Value: 7}
		l.InsertAfter(n, 7)

		for i, n := range []int{1, 2, 3, 5} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should be noop if list is empty", func(t *testing.T) {
		l := New[int]()
		n := &Node[int]{Value: 7}
		l.InsertAfter(n, 7)
		assert.Nil(t, l.head)
	})
}

func TestLinkedList_InsertAt(t *testing.T) {

	t.Run("should insert value at specified index", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 5} {
			l.Append(n)
		}

		l.InsertAt(2, 20)

		for i, n := range []int{1, 2, 20, 3, 5} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should insert value as head if index is 0", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 5} {
			l.Append(n)
		}

		l.InsertAt(0, 20)
		assert.Equal(t, 20, l.head.Value)

		for i, n := range []int{20, 1, 2, 3, 5} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should insert value as second last if index is last", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 5} {
			l.Append(n)
		}

		l.InsertAt(3, 20)

		for i, n := range []int{1, 2, 3, 20, 5} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should be noop if index is out of range", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 5} {
			l.Append(n)
		}

		l.InsertAt(5, 20)

		for i, n := range []int{1, 2, 3, 5} {
			v, ok := l.NthValue(uint(i))
			assert.True(t, ok)
			assert.Equal(t, n, v)
		}
	})

	t.Run("should be noop if list is empty and index is not 0", func(t *testing.T) {
		l := New[int]()
		l.InsertAt(2, 20)
		assert.Nil(t, l.head)
	})

	t.Run("should insert value as head if list is empty and index is 0", func(t *testing.T) {
		l := New[int]()
		l.InsertAt(0, 20)
		assert.Equal(t, 20, l.head.Value)
		n, ok := l.NthValue(0)
		assert.True(t, ok)
		assert.Equal(t, 20, n)
	})
}

func TestLinkedList_Reverse(t *testing.T) {

	t.Run("should reverse the order of nodes in the list", func(t *testing.T) {
		l := New[int]()
		for _, n := range []int{1, 2, 3, 4, 5} {
			l.Append(n)
		}

		l.Reverse()

		assert.Equal(t, 5, l.head.Value)

		expect := []int{5, 4, 3, 2, 1}
		for i, v := range expect {
			n := l.NthNode(uint(i))
			assert.NotNil(t, n)
			assert.Equal(t, v, n.Value)
			if i+1 > len(expect)-1 == false {
				assert.Equal(t, expect[i+1], n.next.Value)
			}
		}
	})
}
