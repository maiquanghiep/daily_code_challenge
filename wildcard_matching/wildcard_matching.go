package wildcardmatching

/*
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).
*/

func isMatch(s string, p string) bool {
	curStrPos := 0
	curPatPos := 0
	curPatAfterStar := -1
	curStrAfterStar := -1

	for curStrPos < len(s) {
		if curPatPos < len(p) && (s[curStrPos] == p[curPatPos] || p[curPatPos] == '?') {
			curPatPos++
			curStrPos++
		} else if curPatPos < len(p) && p[curPatPos] == '*' {
			curPatPos++
			curPatAfterStar = curPatPos
			curStrAfterStar = curStrPos
		} else {
			if curPatAfterStar == -1 {
				return false
			}
			curPatPos = curPatAfterStar + 1
			curStrAfterStar++
			curStrPos = curStrAfterStar
		}
	}

	for curPatPos < len(p) && p[curPatPos] == '*' {
		curPatPos++
	}
	return curPatPos == len(p)
}
