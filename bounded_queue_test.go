package queue

import (
	"sync"
	"testing"
)

func TestBoundedQueue(t *testing.T) {
	count := 100

	q := NewBoundedQueue(10)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < count; i++ {
			q.Enqueue(i)
		}
	}()

	token := 0
	go func() {
		defer wg.Done()

		for i := 0; i < count; i++ {
			v := q.Dequeue()
			if v == nil {
				t.Errorf("got a nil value")
			}
			if v.(int) != i {
				t.Errorf("expect %d but got %v", i, v)
			}

			token++
		}
	}()

	wg.Wait()

	if token != count {
		t.Errorf("expected taken %d but got %d", count, token)
	}

	if q.Len() != 0 {
		t.Errorf("still has %d elements have not been dequeued", q.Len())
	}
}
