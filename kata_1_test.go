package main

import (
	"fmt"
	"testing"
)

func test(t *testing.T) {
	fmt.Print("Test Start")
	Item := NewStack()
	Item.Push(1)
	val, _ := Item.Peek()
	if val != 1 {
		t.Error("Wanted: ", val)
	}
}
