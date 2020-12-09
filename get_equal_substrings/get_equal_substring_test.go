package getequalsubstrings

import "testing"

func TestEqualSubString(t *testing.T) {
	s1 := "abcd"
	s2 := "bcdf"
	maxCost := 3

	result := equalSubstring(s1, s2, maxCost)
	if result != 3 {
		t.Errorf("Expected 3, but got %d", result)
	}

	s3 := "abcd"
	s4 := "cdef"
	maxCost2 := 3

	result2 := equalSubstring(s3, s4, maxCost2)
	if result2 != 1 {
		t.Errorf("Expected 1, but got %d", result2)
	}

	s5 := "abcd"
	s6 := "acde"
	maxCost3 := 0

	result3 := equalSubstring(s5, s6, maxCost3)
	if result3 != 1 {
		t.Errorf("Expected 1, but got %d", result3)
	}

}
