// Package genq provides a generic FIFO queue.
package genq

// Message is a message from the queue
type Message[T any] struct {
	next, prev *Message[T]
	queue      *Queue[T]
	Value      T
}

type Queue[T any] struct {
	root Message[T]
	len  int
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (q *Queue[T]) Len() int { return q.len }

// Init initializes or clears list l.
func (q *Queue[T]) Init() *Queue[T] {
	q.root.next = &q.root
	q.root.prev = &q.root
	q.len = 0
	return q
}

// New returns an initialized list.
func New[T any]() *Queue[T] {
	return new(Queue[T]).Init()
}
