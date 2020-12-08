package longestprefix

import "testing"

func TestLongestPrefix(t *testing.T) {
	strs := []string{
		"flower",
		"flow",
		"flight",
	}

	result := longestCommonPrefix(strs)
	if result != "fl" {
		t.Errorf("Expected fl, but got %s", result)
	}

	strs1 := []string{
		"dog",
		"racecar",
		"car",
	}

	result1 := longestCommonPrefix(strs1)
	if result1 != "" {
		t.Errorf("Expected , but got %s", result1)
	}
}
