package sellstock

import "testing"

func TestSellStock1(t *testing.T) {
	n := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit1(n)
	if result != 5 {
		t.Errorf("Expected 5 but got %d", result)
	}

	n1 := []int{7, 6, 4, 3, 1}
	result1 := maxProfit1(n1)
	if result1 != 0 {
		t.Errorf("Expected 0 but got %d", result1)
	}
}

func TestSellStock2(t *testing.T) {
	n := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit2(n)
	if result != 7 {
		t.Errorf("Expected 7 but got %d", result)
	}

	n1 := []int{7, 6, 4, 3, 1}
	result1 := maxProfit2(n1)
	if result1 != 0 {
		t.Errorf("Expected 0 but got %d", result1)
	}

	n2 := []int{1, 2, 3, 4, 5}
	result2 := maxProfit2(n2)
	if result2 != 4 {
		t.Errorf("Expected 4 but got %d", result1)
	}
}
