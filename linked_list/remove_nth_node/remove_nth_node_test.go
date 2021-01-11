package removenthnode

import (
	linkedlist "daily_code_challenge/lib"
	"testing"
)

func TestMerge(t *testing.T) {
	l1 := linkedlist.NewList(0)
	l1.AddTail(1)
	l1.AddTail(2)
	l1.AddTail(3)
	l1.AddTail(4)
	l1.AddTail(5)

	n := 6
	c := removeNthFromEnd(l1.Head, n)

	expectedResult := "1->2->3->4->5"
	actualResult := c.TraverseNodeToString()
	if actualResult != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, actualResult)
	}

}
