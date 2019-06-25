package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()

	if !q.IsEmpty() ||
		q.size != 0 ||
		q.Size() != 0 {
		t.Error()
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	if q.values[0] != 1 ||
		q.values[1] != 2 ||
		q.values[2] != 3 {
		t.Error()
	}
	if q.Size() != 3 {
		t.Error()
	}

	v := q.Pop()
	i, ok := v.(int)
	if !ok {
		t.Error()
	}
	if i != 1 {
		t.Error()
	}
	if q.Size() != 2 {
		t.Error()
	}

	q.Pop()
	q.Pop()
	q.Pop()
}
