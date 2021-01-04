package textjustificationdp

import "testing"

func TestMaxCost(t *testing.T) {
	s1 := "Geeks for Geeks presents word wrap problem"
	w1 := 15
	r1 := minCost(s1, w1)
	if r1 != 13 {
		t.Errorf("Expected 13 but got %d", r1)
	}

	s2 := "Tushar Roy likes to code"
	w2 := 10
	r2 := minCost(s2, w2)
	if r2 != 26 {
		t.Errorf("Expected 26 but got %d", r2)
	}
}
