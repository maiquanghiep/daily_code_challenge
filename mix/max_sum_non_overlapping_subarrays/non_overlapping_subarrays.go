package nonoverlappingsubarrays

/*
In a given array nums of positive integers, find three non-overlapping subarrays with maximum sum.

Each subarray will be of size k, and we want to maximize the sum of all 3*k entries.

Return the result as a list of indices representing the starting position of each interval (0-indexed). If there are multiple answers, return the lexicographically smallest one.
*/

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	w := make([]int, 0)
	currSum := 0
	for i, v := range nums {
		currSum += v
		if i >= k {
			currSum -= nums[i-k]
		}
		if i >= k-1 {
			w = append(w, currSum)
		}
	}

	left := make([]int, len(w))
	best := 0
	for i := range left {
		if w[i] > w[best] {
			best = i
		}
		left[i] = best
	}

	right := make([]int, len(w))
	best = len(w) - 1
	for i := len(w) - 1; i >= 0; i-- {
		if w[i] >= w[best] {
			best = i
		}
		right[i] = best
	}
	ans := []int{-1, -1, -1}

	for j := k; j < len(w)-k; j++ {
		i, l := left[j-k], right[j+k]
		if ans[0] == -1 || w[i]+w[j]+w[l] > w[ans[0]]+w[ans[1]]+w[ans[2]] {
			ans[0] = i
			ans[1] = j
			ans[2] = l
		}
	}

	return ans
}
