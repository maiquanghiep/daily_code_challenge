package linkedlist

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
}

type List struct {
	head *Node
}

func NewList(v int) *List {
	n := Node{val: v}
	return &List{head: &n}
}

func (l *List) String() string {
	if l.head == nil {
		return fmt.Sprintf("")
	}
	current := l.head
	result := make([]string, 0)
	for {
		text := strconv.Itoa(current.val)
		result = append(result, text)
		if current.next == nil {
			break
		}
		current = current.next
	}

	return fmt.Sprintf(strings.Join(result, "->"))
}

func (l *List) toString() string {
	var sb strings.Builder

	if l.head != nil {
		current := l.head
		for current != nil {
			text := strconv.Itoa(current.val)
			sb.WriteString(text)
			current = current.next
		}
	}

	return sb.String()
}

func (l *List) AddFront(v int) {
	n := Node{val: v}
	if l.head == nil {
		l.head = &n
	} else {
		n.next = l.head
		l.head = &n
	}
}

func (l *List) AddTail(v int) {
	n := Node{val: v}
	if l.head == nil {
		l.head = &n
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = &n
	}
}

func (l *List) Reverse() error {
	if l.head == nil {
		return fmt.Errorf("TranverseError: List is empty")
	}
	var prev *Node
	current := l.head
	for current != nil {
		nextNode := current.next
		current.next = prev
		prev = current
		current = nextNode
	}
	l.head = prev
	return nil
}
