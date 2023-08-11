package main

import "fmt"

// struct -> table with only one column, each can use different types
type Stack struct {
	Stack []int
}

// function with a pointer return (allows to ceate more structs)
func NewStack() *Stack {
	stack := Stack{make([]int, 0)}
	return &stack
}

// function adds a new input at the top of the stack (or bottom of the table) LIFO-rule
func (s *Stack) Push(val int) {
	s.Stack = append(s.Stack, val)
}

// function returns the latest index's element and delets the latest index afterwards
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() == false {
		i := len(s.Stack) - 1
		e := s.Stack[i]
		s.Stack = s.Stack[:i]
		return e, true
	} else {
		return 0, false
	}
}

// function returns the latest index's element
func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() == false {
		i := len(s.Stack) - 1
		e := s.Stack[i]
		return e, true
	} else {
		return 0, false
	}
}

// function indicates whether the input is empty
func (s *Stack) IsEmpty() bool {
	if len(s.Stack) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	stack := NewStack()
	stack.Push(5)
	stack.Push(10)
	stack.Push(15)
	val, _ := stack.Peek()
	fmt.Println(val)
	val, _ = stack.Pop()
	fmt.Println(val)
	isEmpty := stack.IsEmpty()
	fmt.Println(isEmpty)
}
