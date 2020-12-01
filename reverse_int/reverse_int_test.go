package reverseint

import "testing"

func TestReverse(t *testing.T) {

	n := reverse(5)
	if n != 5 {
		t.Errorf("Expected 5, but got %d", n)
	}

	n = reverse(12345)
	if n != 54321 {
		t.Errorf("Expected 54321, but got %d", n)
	}

	n = reverse(-12345)
	if n != -54321 {
		t.Errorf("Expected -54321, but got %d", n)
	}

	// 2147483647 is max int, so 2147483651 must return 0
	n = reverse(1563847412)
	if n != 0 {
		t.Errorf("Expected 0, but got %d", n)
	}

	// -2147483648 is min int, so -1563847412 must return 0
	n = reverse(-1563847412)
	if n != 0 {
		t.Errorf("Expected 0, but got %d", n)
	}
}
