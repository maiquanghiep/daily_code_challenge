package palindromell

import "testing"

func TestPalindromeNumber(t *testing.T) {
	n := -1

	c := isPalindromeNumber(n)
	if c {
		t.Errorf("Expected False, but got %t", c)
	}

	n = 5
	c = isPalindromeNumber(n)
	if !c {
		t.Errorf("Expected True, but got %t", c)
	}

	n = 1000021
	c = isPalindromeNumber(n)
	if c {
		t.Errorf("Expected False, but got %t", c)
	}

	n = 121
	c = isPalindromeNumber(n)
	if !c {
		t.Errorf("Expected True, but got %t", c)
	}

	n = 123456654321
	c = isPalindromeNumber(n)
	if !c {
		t.Errorf("Expected True, but got %t", c)
	}

	n = 1234567890
	c = isPalindromeNumber(n)
	if c {
		t.Errorf("Expected False, but got %t", c)
	}
}
