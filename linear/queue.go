package linear 

import "errors"

var (
  ErrQueueUnderflow = errors.New("queue underflow")
  ErrQueueOverflow = errors.New("queue overflow")
)

type Queue[T any] interface {
  Enqueue(T) error
  Dequeue() (*T, error)
  Peek() (*T, error)
  Size() int
}

type queue[T any] struct {
  items []T
  front, rear int
  capacity int
}

func NewQueue[T any](capacity int) Queue[T] {
  return &queue[T] {
    items: make([]T, capacity),
    front: -1,
    rear: -1,
    capacity: capacity,
  }
}

func (que *queue[T]) Enqueue(t T) error {

  if que.rear >= que.capacity - 1 {
    return ErrQueueOverflow
  }

  que.rear += 1
  que.items[que.rear] = t

  return nil
}

func (que *queue[T]) Dequeue() (*T, error) {

  if que.front >= que.rear {
    return nil, ErrQueueUnderflow
  }

  que.front += 1
  v := que.items[que.front]

  return &v, nil
}

func (que *queue[T]) Peek() (*T, error) {

  if que.front >= que.rear {
    return nil, ErrQueueUnderflow
  }

  v := que.items[que.front+1]

  return &v, nil
}

func (que *queue[T]) Size() int {
  return que.rear - que.front
}
