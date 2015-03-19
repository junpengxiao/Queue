package queue

import "testing"

func TestAll(t *testing.T) {
	q := NewQueue(5)
	i:=0
	for ;; {
		if q.Push(1) != ErrQueueFull {
			i++
		} else {
			break
		}
	}
	if i!=5 {
		t.Errorf("Queue Length isn't right ",i)
	}
	x,err := q.Pop()
	v := x.(int)
	if v != 1 {
		t.Errorf("Value Fetched isn't right")
	}
	if err != nil {
		t.Errorf("Error ", err)
	}
	if q.Push(2) != nil {
		t.Errorf("Push 2, can't do it")
	}
	if q.Push(2) != ErrQueueFull {
		t.Errorf("Push 2 into a full queue")
	}
}

	
