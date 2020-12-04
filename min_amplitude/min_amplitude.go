package minamplitude

import (
	"sort"
)

/*Given an Array A, find the minimum amplitude you can get after changing up to 3 elements.
Amplitude is the range of the array (basically difference between largest and smallest element).
Input: [-1, 3, -1, 8, 5 4]
Output: 2
Explanation: we can change -1, -1, 8 to 3, 4 or 5

Input: [10, 10, 3, 4, 10]
Output: 0
Explanation: change 3 and 4 to 10
*/

/*
The solution is to find the 4 bigest number and 4 smallest number then compare 4 cases.
Complexity is O(n) and space is O(1)
I spent lots of time just to loop through the slice and replace the position
*/
func mintOf(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minAmplitude(n []int) int {
	max := []int{n[0], n[1], n[2], n[3]}
	sort.Slice(max, func(i, j int) bool {
		return max[i] > max[j]
	})

	min := []int{max[3], max[2], max[1], max[0]}
	for i := 4; i < len(n); i++ {
		for j, e := range max {
			if n[i] > e {
				for k := len(max) - 1; k > j; k-- {
					max[k] = max[k-1]
				}
				max[j] = n[i]
				break
			}
		}
		for j, e := range min {
			if n[i] < e {
				for k := len(min) - 1; k > j; k-- {
					min[k] = min[k-1]
				}
				min[j] = n[i]
				break
			}
		}
	}

	// Replace 3 bigest numbers
	result := max[3] - min[0]

	// Replace 2 bigest numbers, 1 smallest numbers
	result = mintOf(result, max[2]-min[1])

	// Replace 1 bigest number, 2 smallest numbers
	result = mintOf(result, max[1]-min[2])

	// Replace 3 smallest number
	result = mintOf(result, max[0]-min[3])

	return result
}
