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
func (q *Queue[T]) insert(m, at *Message[T]) *Message[T] {
	m.prev = at
	m.next = at.next
	m.prev.next = m
	m.next.prev = m
	m.queue = q
	q.len++
	return m
}

// insertValue is a convenience wrapper for insert(&Element{Value: v}, at).
func (q *Queue[T]) insertValue(v T, at *Message[T]) *Message[T] {
	return q.insert(&Message[T]{Value: v}, at)
}

// remove removes m from its queue, decrements q.len
func (q *Queue[T]) remove(m *Message[T]) {
	m.prev.next = m.next
	m.next.prev = m.prev
	m.next = nil // avoid memory leaks
	m.prev = nil // avoid memory leaks
	m.queue = nil
	q.len--
}

// Enqueue adds a new Message[T] m with value v at the back of Queue[T] q and returns m.
func (q *Queue[T]) Enqueue(v T) *Message[T] {
	q.lazyInit()
	return q.insertValue(v, q.root.prev)
}

// Dequeue returns the next Message[T] and removes it from Queue[T] q or nil if the queue is empty.
func (q *Queue[T]) Dequeue() *Message[T] {
	if q.len == 0 {
		return nil
	}
	m := q.root.next
	q.remove(m)
	return m
}

// Len returns the number of elements of Queue[T] q.
// The complexity is O(1).
func (q *Queue[T]) Len() int { return q.len }

// Init initializes or clears queue q.
func (q *Queue[T]) Init() *Queue[T] {
	q.root.next = &q.root
	q.root.prev = &q.root
	q.len = 0
	return q
}

// New returns an initialized queue.
func New[T any]() *Queue[T] {
	return new(Queue[T]).Init()
}
