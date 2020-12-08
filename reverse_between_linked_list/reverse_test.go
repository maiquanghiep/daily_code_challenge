package reverse

import (
	linkedlist "daily_code_challenge/lib"
	"testing"
)

func TestMerge(t *testing.T) {
	l1 := linkedlist.NewList(1)
	l1.AddTail(2)
	l1.AddTail(3)
	l1.AddTail(4)
	l1.AddTail(5)

	m := 2
	n := 4

	c := reverseBetween(l1.Head, m, n)
	expectedResult := "1->4->3->2->5"
	actualResult := c.TraverseNodeToString()
	if actualResult != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, actualResult)
	}

	l2 := linkedlist.NewList(1)
	l2.AddTail(2)

	m = 1
	n = 2
	c = reverseBetween(l2.Head, m, n)
	expectedResult = "2->1"
	actualResult = c.TraverseNodeToString()
	if actualResult != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, actualResult)
	}

}
