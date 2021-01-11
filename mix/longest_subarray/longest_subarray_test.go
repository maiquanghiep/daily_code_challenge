package longestsubarray

import "testing"

func TestLongestSubArray(t *testing.T) {
	nums := []int{8, 2, 4, 7}
	limit := 4
	r := longestSubarray(nums, limit)
	if r != 2 {
		t.Errorf("Expected 2, but got %d", r)
	}

	nums = []int{10, 1, 2, 4, 7, 2}
	limit = 5
	r = longestSubarray(nums, limit)
	if r != 4 {
		t.Errorf("Expected 4, but got %d", r)
	}

	nums = []int{4, 2, 2, 2, 4, 4, 2, 2}
	limit = 0
	r = longestSubarray(nums, limit)
	if r != 3 {
		t.Errorf("Expected 3, but got %d", r)
	}

}
