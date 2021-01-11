package nonoverlappingsubarrays

import "testing"

func TestMaxSumOfThreeSubArrays(t *testing.T) {
	nums := []int{4, 5, 10, 6, 11, 17, 4, 11, 1, 3}
	k := 1
	result := maxSumOfThreeSubarrays(nums, k)
	expectedResult := []int{4, 5, 7}
	for i, v := range result {
		if v != expectedResult[i] {
			t.Errorf("Expected %d but got %d", expectedResult[i], v)
		}
	}

	nums2 := []int{7, 13, 20, 19, 19, 2, 10, 1, 1, 19}
	k2 := 3
	result2 := maxSumOfThreeSubarrays(nums2, k2)

	expectedResult2 := []int{1, 4, 7}
	for i, v := range result2 {
		if v != expectedResult2[i] {
			t.Errorf("Expected %d but got %d", expectedResult2[i], v)
		}
	}

}
