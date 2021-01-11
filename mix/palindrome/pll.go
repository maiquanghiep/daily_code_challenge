package palindromell

import (
	linkedlist "daily_code_challenge/lib"
)

/*
	Given a singly linked list, determine if it is a palindrome.

	Example 1:

	Input: 1->2
	Output: false
	Example 2:

	Input: 1->2->2->1
	Output: true
	Follow up:
	Could you do it in O(n) time and O(1) space?
*/

/*
	An easy solution with O(n) time an O(n) space is store the linked list in
	an array and check for palindrome
	-> The solution in O(1) space would be go to the middle of the linked list
	by using 2 pointer (One go to Next node each time and another go 2 nodes),
	if the linked list is odd (The fast pointer point to the Next node is Null)
	we can return False
	Otherwise, we can reverse the first or the second part of the linked list and
	compare them.
*/

func isPalindrome(head *linkedlist.Node) bool {
	// If only one node, return false
	if head == nil || head.Next == nil {
		return true
	}

	slowTraverse := head.Next
	fastTraverse := head.Next.Next

	// There is only 2 nodes so compare them
	if fastTraverse == nil {
		if head.Val == slowTraverse.Val {
			return true
		}
		return false
	}

	var right *linkedlist.Node
	for {
		// odd
		if fastTraverse.Next == nil {
			right = &linkedlist.Node{Val: slowTraverse.Val, Next: slowTraverse.Next}
			slowTraverse.Next = nil
			break
		}
		// even
		if fastTraverse.Next.Next == nil {
			right = slowTraverse.Next
			slowTraverse.Next = nil
			break
		}

		slowTraverse = slowTraverse.Next
		fastTraverse = fastTraverse.Next.Next
	}

	var prev *linkedlist.Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	cur = prev

	for cur != nil {
		if cur.Val != right.Val {
			return false
		}
		cur = cur.Next
		right = right.Next
	}

	return true
}
