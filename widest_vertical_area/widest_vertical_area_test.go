package widestverticalarea

import (
	"testing"
)

func TestWidestVerticalArea(t *testing.T) {
	l := [][]int{
		{8, 7},
		{9, 9},
		{7, 4},
		{9, 7},
	}
	m := maxWidthOfVerticalArea(l)
	if m != 1 {
		t.Errorf("Expected 1, but got %d", m)
	}

	l1 := [][]int{
		{3, 1},
		{9, 0},
		{1, 0},
		{1, 4},
		{5, 3},
		{8, 8},
	}
	m = maxWidthOfVerticalArea(l1)

	if m != 3 {
		t.Errorf("Expected 3, but got %d", m)
	}
}
