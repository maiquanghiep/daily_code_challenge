package slidingwindowmaximum

/*
You are given an array of integers nums, there is a sliding window of size k
which is moving from the very left of the array to the very right.
You can only see the k numbers in the window. Each time the sliding window moves right by one position.

Return the max sliding window.
*/
type deque struct {
	elements []int
}

func (d *deque) push(e int) {
	d.elements = append(d.elements, e)
}

func (d *deque) getFist() int {
	return d.elements[0]
}

func (d *deque) popFirst() {
	d.elements = d.elements[1:]
}

func (d *deque) getLast() int {
	return d.elements[len(d.elements)-1]
}

func (d *deque) popLast() {
	d.elements = d.elements[:len(d.elements)-1]
}

func (d *deque) empty() bool {
	return len(d.elements) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	dq := &deque{}
	ans := []int{}
	for i, v := range nums {
		if i >= k && (dq.getFist() == nums[i-k]) {
			dq.popFirst()
		}

		for dq.empty() == false && v > dq.getLast() {
			dq.popLast()
		}

		dq.push(v)

		if i >= k-1 {
			ans = append(ans, dq.getFist())
		}

	}
	return ans
}
