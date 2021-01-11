package longestprefix

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Input: strs = ["flower","flow","flight"]
Output: "fl"

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
*/
func findCommonPrefix(s1 string, s2 string, maxCommonPrefix int) int {
	if len(s2) < maxCommonPrefix {
		maxCommonPrefix = len(s2)
	}
	for maxCommonPrefix > 0 {
		if s1[:maxCommonPrefix] == s2[:maxCommonPrefix] {
			return maxCommonPrefix
		}
		maxCommonPrefix--
	}
	return maxCommonPrefix
}
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	maxCommonPrefix := len(strs[0])

	for i := 1; i < len(strs); i++ {
		maxCommonPrefix = findCommonPrefix(strs[0], strs[i], maxCommonPrefix)
		if maxCommonPrefix == 0 {
			return ""
		}
	}

	return strs[0][:maxCommonPrefix]
}
