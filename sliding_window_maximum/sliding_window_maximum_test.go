package slidingwindowmaximum

import (
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	r1 := maxSlidingWindow(nums1, k1)
	actualResult := []int{3, 3, 5, 5, 6, 7}
	for i, v := range actualResult {
		if v != r1[i] {
			t.Errorf("Expected %d, but got %d", actualResult[i], v)
			return
		}
	}

	nums2 := []int{9, 10, 9, -7, -4, -8, 2, -6}
	k2 := 5
	r2 := maxSlidingWindow(nums2, k2)
	actualResult2 := []int{10, 10, 9, 2}
	if len(r2) != len(actualResult2) {
		t.Errorf("Expected %d, but got %d", actualResult2, r2)
		return
	}
	for i, v := range actualResult2 {
		if v != r2[i] {
			t.Errorf("Expected %d, but got %d", actualResult2[i], v)
			return
		}
	}
}
