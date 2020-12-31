package textjustificationdp

import (
	"strings"
)

/*
Given a sequence of words, and a limit on the number of characters that can be put in one line (line width). Put line breaks in the given sequence such that the lines are printed neatly. Assume that the length of each word is smaller than the line width.

The word processors like MS Word do task of placing line breaks. The idea is to have balanced lines. In other words, not have few lines with lots of extra spaces and some lines with small amount of extra spaces.
*/
func calculateCost(c, w int) int {
	cost := w - c
	if cost < 0 {
		return -1
	}
	return cost * cost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(s string, w int) int {
	texts := strings.Split(s, " ")
	dp := make([][]int, len(texts))
	for i := range dp {
		dp[i] = make([]int, len(texts))
	}

	for i := range dp {
		count := 0
		for j := i; j < len(texts); j++ {
			count += len(texts[j]) + j - i
			dp[i][j] = calculateCost(count, w)
		}
	}

	for i := len(texts) - 1; i >= 0; i-- {
		minCost := dp[i][i]
		if minCost == -1 {
			for j := len(texts); j >= i; j-- {
				if dp[i][j] == -1 {
					continue
				}
				minCost = min(minCost, dp[i][j] + )
			}
		}
	}

	return w
}
