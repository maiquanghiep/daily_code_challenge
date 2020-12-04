package minamplitude

import "testing"

func TestMinAmplitude(t *testing.T) {
	n := []int{-1, 3, -1, 8, 5, 4}
	r := minAmplitude(n)
	if r != 2 {
		t.Errorf("Expected 2, but got %d", r)
	}

	n1 := []int{10, 10, 3, 4, 10}
	r = minAmplitude(n1)
	if r != 0 {
		t.Errorf("Expected 0, but got %d", r)
	}
}
