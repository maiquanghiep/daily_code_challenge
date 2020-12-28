package maximumsubarray

import "testing"

func TestMaximumSubarray(t *testing.T) {
	nums1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maximum1 := maxSubArray(nums1)

	if maximum1 != 6 {
		t.Errorf("Expetec 6 but got %d", maximum1)
	}

	nums2 := []int{-1}
	maximum2 := maxSubArray(nums2)

	if maximum2 != -1 {
		t.Errorf("Expetec -1 but got %d", maximum2)
	}
}
