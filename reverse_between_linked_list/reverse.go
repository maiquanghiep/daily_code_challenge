package reverse

import (
	linkedlist "daily_code_challenge/lib"
)

/*
Reverse a linked list from position m to n. Do it in one-pass.
Note: 1 ≤ m ≤ n ≤ length of list.
Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
*/

func reverseBetween(head *linkedlist.Node, m int, n int) *linkedlist.Node {
	cur := head

	count := 0
	var start, end, before *linkedlist.Node

	if m == 1 {
		start = head
	}
	for cur != nil {
		count++
		if count == m-1 {
			before = cur
			start = cur.Next
		}
		if count == n {
			end = cur
		}
		cur = cur.Next
	}

	if count == 1 || (m == n) {
		return head
	}

	prev := end.Next
	for i := 0; i < n-m+1; i++ {
		nextNode := start.Next
		start.Next = prev
		prev = start
		start = nextNode
	}
	if before != nil {
		before.Next = prev
	} else {
		head = prev
	}

	return head
}
