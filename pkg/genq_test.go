package genq

import "testing"

func checkQueueLen[T any](t *testing.T, q *Queue[T], len int) bool {
	if n := q.Len(); n != len {
		t.Errorf("q.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func TestQueue(t *testing.T) {
	q := New[bool]()
	checkQueueLen[bool](t, q, 0)
}
