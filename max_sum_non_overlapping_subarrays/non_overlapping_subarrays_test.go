package nonoverlappingsubarrays

import "testing"

func TestMaxSumOfThreeSubArrays(t *testing.T) {
	nums := []int{7, 13, 20, 19, 19, 2, 10, 1, 1, 19}
	k := 3
	result := maxSumOfThreeSubarrays(nums, k)
	expectedResult := []int{0, 3, 5}
	for i, v := range result {
		if v != expectedResult[i] {
			t.Errorf("Expected %d but got %d", expectedResult[i], v)
		}
	}
}
