package wildcardmatching

import "testing"

func TestWildCardMatching(t *testing.T) {
	s1 := "zacabz"
	p1 := "*a?b*"
	isMatch1 := isMatch(s1, p1)
	if isMatch1 {
		t.Errorf("Expected false but got %t", isMatch1)
	}

	s2 := "aa"
	p2 := "a"
	isMatch2 := isMatch(s2, p2)
	if isMatch2 {
		t.Errorf("Expected false but got %t", isMatch2)
	}

	s3 := "abcabczzzde"
	p3 := "*abc???de*"
	isMatch3 := isMatch(s3, p3)
	if !isMatch3 {
		t.Errorf("Expected true but got %t", isMatch3)
	}

	s4 := "adceb"
	p4 := "*a*b"
	isMatch4 := isMatch(s4, p4)
	if !isMatch4 {
		t.Errorf("Expected true but got %t", isMatch4)
	}

	s5 := "abcdcb"
	p5 := "a*c?b"
	isMatch5 := isMatch(s5, p5)
	if isMatch5 {
		t.Errorf("Expected false but got %t", isMatch5)
	}

}
