package linkedlist

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

type List struct {
	Head *Node
}

func NewList(v int) *List {
	n := Node{Val: v}
	return &List{Head: &n}
}

func (l *List) String() string {
	if l.Head == nil {
		return fmt.Sprintf("")
	}
	current := l.Head
	result := make([]string, 0)
	for {
		text := strconv.Itoa(current.Val)
		result = append(result, text)
		if current.Next == nil {
			break
		}
		current = current.Next
	}

	return fmt.Sprintf(strings.Join(result, "->"))
}

func (l *List) toString() string {
	var sb strings.Builder

	if l.Head != nil {
		current := l.Head
		for current != nil {
			text := strconv.Itoa(current.Val)
			sb.WriteString(text)
			current = current.Next
		}
	}

	return sb.String()
}

func (l *List) AddFront(v int) {
	n := Node{Val: v}
	if l.Head == nil {
		l.Head = &n
	} else {
		n.Next = l.Head
		l.Head = &n
	}
}

func (l *List) AddTail(v int) {
	n := Node{Val: v}
	if l.Head == nil {
		l.Head = &n
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &n
	}
}

func (l *List) Reverse() error {
	if l.Head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	var prev *Node
	current := l.Head
	for current != nil {
		NextNode := current.Next
		current.Next = prev
		prev = current
		current = NextNode
	}
	l.Head = prev
	return nil
}

func (n *Node) TraverseNodeToString() string {
	cur := n
	vals := make([]string, 0)
	for cur != nil {
		vals = append(vals, strconv.Itoa(cur.Val))
		cur = cur.Next
	}

	var sb strings.Builder
	sb.WriteString(strings.Join(vals, "->"))

	return sb.String()
}
