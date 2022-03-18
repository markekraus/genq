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

// lazyInit lazily initializes a zero queue value.
func (q *Queue[T]) lazyInit() {
	if q.root.next == nil {
		q.Init()
	}
}

// insert inserts e after at, increments l.len, and returns e.
func (q *Queue[T]) insert(e, at *Message[T]) *Message[T] {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.queue = q
	q.len++
	return e
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (q *Queue[T]) insertValue(v T, at *Message[T]) *Message[T] {
	return q.insert(&Message[T]{Value: v}, at)
}

// Enqueue adds a new Message[T] m with value v at the back of Queue[T] q and returns m.
func (q *Queue[T]) Enqueue(v T) *Message[T] {
	q.lazyInit()
	return q.insertValue(v, q.root.prev)
}

// Len returns the number of elements of Queue[T] q.
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
