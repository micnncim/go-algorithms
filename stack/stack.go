package stack

import (
	"sync"
)

type Stack struct {
	values []interface{}
	size   int
	mu     *sync.Mutex
}

func New() *Stack {
	s := new(Stack)
	s.values = make([]interface{}, 0)
	s.mu = new(sync.Mutex)
	return s
}

func (s *Stack) Size() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.size
}

func (s *Stack) IsEmpty() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.size == 0
}

func (s *Stack) Pop() interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.size == 0 {
		return nil
	}
	v := s.values[s.size-1]
	if s.size == 1 {
		s.values = make([]interface{}, 0)
	} else {
		s.values = s.values[:s.size-1]
	}
	s.size--
	return v
}

func (s *Stack) Push(v interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values = append(s.values, v)
	s.size++
}
