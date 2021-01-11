package mergelinkedlist

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

	l2 := linkedlist.NewList(1000000)
	l2.AddTail(1000001)
	l2.AddTail(1000002)

	a := 3
	b := 4

	c := mergeInBetween(l1.Head, a, b, l2.Head)
	expectedResult := "0->1->2->1000000->1000001->1000002->5"
	actualResult := c.TraverseNodeToString()
	if actualResult != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, actualResult)
	}

	l3 := linkedlist.NewList(0)
	l3.AddTail(1)
	l3.AddTail(2)
	l3.AddTail(3)
	l3.AddTail(4)
	l3.AddTail(5)
	l3.AddTail(6)

	l4 := linkedlist.NewList(1000000)
	l4.AddTail(1000001)
	l4.AddTail(1000002)
	l4.AddTail(1000003)
	l4.AddTail(1000004)

	a = 2
	b = 5
	c1 := mergeInBetween(l3.Head, a, b, l4.Head)
	expectedResult = "0->1->1000000->1000001->1000002->1000003->1000004->6"
	actualResult = c1.TraverseNodeToString()
	if actualResult != expectedResult {
		t.Errorf("Expected %s, but got %s", expectedResult, actualResult)
	}
}
