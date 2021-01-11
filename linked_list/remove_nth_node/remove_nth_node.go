package removenthnode

import (
	linkedlist "daily_code_challenge/lib"
)

/*
Given the head of a linked list, remove the nth node from the end of
the list and return its head.
Follow up: Could you do this in one pass?
*/
func removeNthFromEnd(head *linkedlist.Node, n int) *linkedlist.Node {
	slow, fast := head, head
	prev := &linkedlist.Node{}
	for count := 1; count < n; count++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		prev = slow
		slow = slow.Next
	}

	if slow == fast && prev.Next == nil {
		return nil
	}
	if prev.Next == nil {
		return slow.Next
	}

	prev.Next = slow.Next

	return head
}
