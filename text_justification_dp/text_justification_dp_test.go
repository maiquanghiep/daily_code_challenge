package textjustificationdp

import "testing"

func TestMaxCost(t *testing.T) {
	s1 := "Geeks for Geeks presents word wrap problem"
	w1 := 15
	r1 := minCost(s1, w1)
	if r1 != 13 {
		t.Errorf("Expected 13 but got %d", r1)
	}
}
