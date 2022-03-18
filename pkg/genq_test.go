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

func TestEnqueue(t *testing.T) {
	q := New[bool]()
	checkQueueLen[bool](t, q, 0)
	q.Enqueue(true)
	checkQueueLen[bool](t, q, 1)
	q.Enqueue(false)
	checkQueueLen[bool](t, q, 2)
}

func TestDequeue(t *testing.T) {
	q := New[bool]()
	q.Enqueue(true)
	q.Enqueue(false)
	q.Enqueue(true)
	q.Enqueue(false)
	checkQueueLen[bool](t, q, 4)
	m := q.Dequeue()
	checkQueueLen[bool](t, q, 3)
	if m.Value != true {
		t.Errorf("m.Value = %v, want %v", m.Value, true)
	}
	m = q.Dequeue()
	checkQueueLen[bool](t, q, 2)
	if m.Value != false {
		t.Errorf("m.Value = %v, want %v", m.Value, false)
	}
	m = q.Dequeue()
	checkQueueLen[bool](t, q, 1)
	if m.Value != true {
		t.Errorf("m.Value = %v, want %v", m.Value, true)
	}
	m = q.Dequeue()
	checkQueueLen[bool](t, q, 0)
	if m.Value != false {
		t.Errorf("m.Value = %v, want %v", m.Value, false)
	}
}
