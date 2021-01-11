package palindromell

import (
	linkedlist "daily_code_challenge/lib"
	"testing"
)

func TestPalindromeSpecialCase(t *testing.T) {
	l := linkedlist.NewList(1)
	l.AddTail(2)

	c := isPalindrome(l.Head)
	if c {
		t.Errorf("Expected False, but got %t", c)
	}

	l1 := linkedlist.NewList(1)
	l1.AddTail(1)

	c = isPalindrome(l1.Head)
	if !c {
		t.Errorf("Expected True, but got %t", c)
	}
}

func TestPalindrome(t *testing.T) {
	l := linkedlist.NewList(1)
	l.AddTail(2)
	l.AddTail(2)
	l.AddTail(1)

	c := isPalindrome(l.Head)
	if !c {
		t.Errorf("Expected True, but got %t", c)
	}

	l.AddTail(0)
	c = isPalindrome(l.Head)
	if c {
		t.Errorf("Expected False, but got %t", c)
	}

	l1 := linkedlist.NewList(1)
	l1.AddTail(2)
	l1.AddTail(3)
	l1.AddTail(2)
	l1.AddTail(1)

	c1 := isPalindrome(l1.Head)
	if !c1 {
		t.Errorf("Expected True, but got %t", c)
	}

	l1.AddTail(0)
	c1 = isPalindrome(l.Head)
	if c {
		t.Errorf("Expected False, but got %t", c)
	}

}
