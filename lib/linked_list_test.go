package linkedlist

import (
	"testing"
)

func TestLinkedListInit(t *testing.T) {
	l := NewList(2)
	l.AddFront(1)
	l.AddTail(3)

	testString := l.toString()
	if testString != "123" {
		t.Errorf("Expected 123, but got %s", testString)
	}

	l.Reverse()
	testString = l.toString()
	if testString != "321" {
		t.Errorf("Expected 321, but got %s", testString)
	}
}
