package getequalsubstrings

/*
You are given two strings s and t of the same length. You want to change s to t. Changing the i-th character of s to i-th character of t costs |s[i] - t[i]| that is, the absolute difference between the ASCII values of the characters.

You are also given an integer maxCost.

Return the maximum length of a substring of s that can be changed to be the same as the corresponding substring of twith a cost less than or equal to maxCost.

If there is no substring from s that can be changed to its corresponding substring from t, return 0.

*/
func changeCost(s byte, t byte) int {
	if s > t {
		return int(s - t)
	}
	return int(t - s)
}

func equalSubstring(s string, t string, maxCost int) int {
	var cur int
	result := 0
	for start, end := 0, 0; end < len(s); end++ {
		cost := changeCost(s[end], t[end])
		cur = cur + cost
		if cur <= maxCost {
			if end-start+1 > result {
				result = end - start + 1
			}
		} else {
			cur = cur - changeCost(s[start], t[start])
			start++
		}
	}

	return result
}
