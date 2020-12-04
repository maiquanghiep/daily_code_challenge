package longestsubarray

func longestSubarray(nums []int, limit int) int {
	var result int
	var i, j int
	min := []int{}
	max := []int{}
	for j < len(nums) {
		for len(min) > 0 && nums[min[len(min)-1]] >= nums[j] {
			min = min[:len(min)-1]
		}
		min = append(min, j)

		for len(max) > 0 && nums[max[len(max)-1]] <= nums[j] {
			max = max[:len(max)-1]
		}
		max = append(max, j)

		if nums[max[0]]-nums[min[0]] > limit {
			i++
			if i > min[0] {
				min = min[1:]
			}
			if i > max[0] {
				max = max[1:]
			}
		} else {
			if j-i+1 > result {
				result = j - i + 1
			}
			j++
		}
	}
	return result
}
