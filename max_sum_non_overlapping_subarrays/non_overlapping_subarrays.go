package nonoverlappingsubarrays

import (
	"strconv"
)

func sumAtPos(nums []int, k int, pos int) int {
	if pos+k > len(nums) {
		return 0
	}

	var result int
	for i := 0; i < k; i++ {
		result += nums[i+pos]
	}
	return result
}

func maxAtPos(nums []int, count int, k int, pos int, memo map[string]int) (int, []int) {
	if pos+k > len(nums) || count == 0 {
		return 0, nil
	}

	key := strconv.Itoa(count) + "," + strconv.Itoa(pos)
	if v, ok := memo[key]; ok {
		return v, nil
	}

	var result int

	array := make([]int, 3)
	for i := 0; i < k; i++ {
		maxAt, resultArray := maxAtPos(nums, count-1, k, pos+k+i, memo)
		if result < sumAtPos(nums, k, pos+i)+maxAt {
			result = sumAtPos(nums, k, pos+i) + maxAt
			copy(array, resultArray)
			array[count-1] = pos + i
		}
	}

	memo[key] = result

	return result, array
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	memo := make(map[string]int)
	_, array := maxAtPos(nums, 3, k, 0, memo)

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
