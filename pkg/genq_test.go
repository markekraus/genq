package genq

import "testing"

type myType struct {
	a, b int
}

func checkQueueLen[T any](t *testing.T, q *Queue[T], len int) bool {
	if n := q.Len(); n != len {
		t.Errorf("q.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func TestQueue(t *testing.T) {
	q := New[bool]()
	checkQueueLen(t, q, 0)
}

func TestEnqueue(t *testing.T) {
	q := New[bool]()
	checkQueueLen(t, q, 0)
	q.Enqueue(true)
	checkQueueLen(t, q, 1)
	q.Enqueue(false)
	checkQueueLen(t, q, 2)
}

func TestDequeue(t *testing.T) {
	q := New[bool]()
	q.Enqueue(true)
	q.Enqueue(false)
	q.Enqueue(true)
	q.Enqueue(false)
	checkQueueLen(t, q, 4)
	m := q.Dequeue()
	checkQueueLen(t, q, 3)
	if m.Value != true {
		t.Errorf("m.Value = %v, want %v", m.Value, true)
	}
	m = q.Dequeue()
	checkQueueLen(t, q, 2)
	if m.Value != false {
		t.Errorf("m.Value = %v, want %v", m.Value, false)
	}
	m = q.Dequeue()
	checkQueueLen(t, q, 1)
	if m.Value != true {
		t.Errorf("m.Value = %v, want %v", m.Value, true)
	}
	m = q.Dequeue()
	checkQueueLen(t, q, 0)
	if m.Value != false {
		t.Errorf("m.Value = %v, want %v", m.Value, false)
	}
}

func TestStruct(t *testing.T) {
	m1 := &myType{1, 2}
	q := New[*myType]()
	q.Enqueue(m1)
	checkQueueLen(t, q, 1)
	m2 := q.Dequeue().Value
	if m2 != m1 {
		t.Errorf("e2 = %v, want %v", m2, m1)
	}
	if m2.a != 1 {
		t.Errorf("e2.a = %v, want %v", m2.a, 1)
	}
	if m2.b != 2 {
		t.Errorf("e2.b = %v, want %v", m2.b, 2)
	}
}
