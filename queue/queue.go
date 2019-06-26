package queue

import (
	"sync"
)

type Queue struct {
	values []interface{}
	size   int
	mu     *sync.Mutex
}

func New() *Queue {
	return &Queue{
		values: make([]interface{}, 0),
		mu:     new(sync.Mutex),
	}
}

func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.size
}

func (q *Queue) IsEmpty() bool {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.size == 0
}

func (q *Queue) Pop() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.size == 0 {
		return nil
	}
	v := q.values[0]
	if q.size == 1 {
		q.values = make([]interface{}, 0)
	} else {
		q.values = q.values[1:]
	}
	q.size--
	return v
}

func (q *Queue) Push(v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.values = append(q.values, v)
	q.size++
}
