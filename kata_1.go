package main

import "fmt"

// struct -> table with only one column, each can use different types
type Item struct {
	Item []interface{}
}

// function with a pointer return (allows to ceate more structs)
func NewStack() *Item {
	Item := Item{make([]interface{}, 0)}
	return &Item
}

// function adds a new input at the top of the Item (or bottom of the table) LIFO-rule
func (s *Item) Push(val interface{}) {
	s.Item = append(s.Item, val)
}

// function returns the latest index's element and delets the latest index afterwards
func (s *Item) Pop() (interface{}, bool) {
	i := 0
	b := false
	if s.IsEmpty() == false {
		i = len(s.Item) - 1
		b = true
	}
	e := s.Item[i]
	s.Item = s.Item[:i]
	return e, b
}

// function returns the latest index's element
func (s *Item) Peek() (interface{}, bool) {
	i := 0
	b := false
	if s.IsEmpty() == false {
		i = len(s.Item) - 1
		b = true
	}
	e := s.Item[i]
	return e, b
}

// function indicates whether the input is empty
func (s *Item) IsEmpty() bool {
	b := false
	if len(s.Item) == 0 {
		b = true
	}
	return b
}

func main() {
	Item := NewStack()
	Item.Push(5)
	Item.Push(10)
	Item.Push(15)
	val, _ := Item.Peek()
	fmt.Println(val)
	val, _ = Item.Pop()
	fmt.Println(val)
	isEmpty := Item.IsEmpty()
	fmt.Println(isEmpty)
}
