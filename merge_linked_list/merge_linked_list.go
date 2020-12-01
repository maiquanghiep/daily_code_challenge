package mergelinkedlist

import (
	linkedlist "daily_code_challenge/lib"
)

/*
You are given two linked lists: list1 and list2 of sizes n and m respectively.

Remove list1's nodes from the ath node to the bth node, and put list2 in their place.
*/
func mergeInBetween(list1 *linkedlist.Node, a int, b int, list2 *linkedlist.Node) *linkedlist.Node {
	cur1 := list1
	cur2 := list2
	for i := 0; i < a-1; i++ {
		cur1 = cur1.Next
	}
	lastNodeFirstLink := cur1.Next
	cur1.Next = cur2

	for i := a; i <= b; i++ {
		lastNodeFirstLink = lastNodeFirstLink.Next
	}

	for {
		if cur2.Next == nil {
			break
		}
		cur2 = cur2.Next
	}
	cur2.Next = lastNodeFirstLink

	return list1
}
