package stack

import "testing"

func TestStack(t *testing.T) {
	s := New()

	if !s.IsEmpty() ||
		s.size != 0 ||
		s.Size() != 0 {
		t.Error()
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.values[0] != 1 ||
		s.values[1] != 2 ||
		s.values[2] != 3 {
		t.Error()
	}
	if s.Size() != 3 {
		t.Error()
	}

	v := s.Pop()
	i, ok := v.(int)
	if !ok {
		t.Error()
	}
	if i != 3 {
		t.Error()
	}
	if s.Size() != 2 {
		t.Error()
	}

	s.Pop()
	s.Pop()
	s.Pop()
}
