package main

import "errors"

var (
  ErrStackOverflow = errors.New("stack overflow")
  ErrStackUnderflow = errors.New("stack underflow")
)

type Stack[T any] interface {
  Push(T) error
  Pop() (*T, error)
  Peek() (*T, error)
  Size() int
}

type stack[T any] struct {
  items []T
  capacity int 
  top int
}

func (stk *stack[T]) Push(t T) error {

  if stk.top >= stk.capacity-1 {
    return ErrStackOverflow    
  }

  stk.items[stk.top+1] = t
  stk.top += 1 

  return nil
}

func (stk *stack[T]) Pop() (*T, error) {
 
  if stk.top < 0 {
    return nil, ErrStackUnderflow 
  }

  item := stk.items[stk.top]
  stk.top -= 1

  return &item, nil
}

func (stk *stack[T]) Peek() (*T, error) {
 
  if stk.top < 0 {
    return nil, ErrStackUnderflow 
  }

  item := stk.items[stk.top]

  return &item, nil
}

func (stk *stack[T]) Size() int {
  
  return stk.top + 1
}

func NewStack[T any](capacity int) Stack[T] {
  return  &stack[T] {
    items: make([]T,capacity),
    capacity: capacity,
    top: -1,
  } 
}

