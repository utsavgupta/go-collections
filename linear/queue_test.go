package linear

import "testing"

func TestEnqueueDequeSuccess(t *testing.T) {


  testTable := []struct{
    size int
    items []int
  }{
    {5, []int{7,2,4,5,2}},
    {1, []int{2}},
  }

  for _, test := range testTable {
    q := NewQueue[int](5)

    for _, item := range test.items {
      q.Enqueue(item)
    } 

    for _, item := range test.items {
      v, err := q.Dequeue()

      if err != nil {
        t.Errorf("Unexpected error %v", err)
      }

      if *v != item {
        t.Errorf("expected %v, got %v", item, *v)
      }
    } 
  }
  
}

func TestQueueOverflow(t *testing.T) {
 
  items := []int{7,2,4,5,2}

  q := NewQueue[int](5)

  for _, item := range items {
    if err := q.Enqueue(item); err != nil {
      t.Errorf("Unexpected error %v", err)
    }
  }

  if err := q.Enqueue(50); err != ErrQueueOverflow {
    t.Errorf("Expected overflow error, but did not receive one.")
  }
  
}

func TestQueueUnderflow(t *testing.T) {
 
  items := []int{7,2,4,5,2}

  q := NewQueue[int](5)

  for _, item := range items {
    if err := q.Enqueue(item); err != nil {
      t.Errorf("Unexpected error %v", err)
    }
  }

  for range items {
    if _, err := q.Dequeue(); err != nil {
      t.Errorf("Unexpected error %v", err)
    }
  }

  if _, err := q.Dequeue(); err != ErrQueueUnderflow {
    t.Errorf("Expected underflow error, but did not receive one.")
  } 
}
