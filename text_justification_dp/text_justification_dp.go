package textjustificationdp

import (
	"math"
	"strings"
)

/*
Given a sequence of words, and a limit on the number of characters that can be put in one line (line width). Put line breaks in the given sequence such that the lines are printed neatly. Assume that the length of each word is smaller than the line width.

The word processors like MS Word do task of placing line breaks. The idea is to have balanced lines. In other words, not have few lines with lots of extra spaces and some lines with small amount of extra spaces.
*/
func calculateCost(c, w int) int {
	cost := w - c
	if cost < 0 {
		return math.MaxInt16
	}
	return cost * cost
}

func minCost(s string, w int) int {
	texts := strings.Split(s, " ")
	dp := make([][]int, len(texts))
	for i := range dp {
		dp[i] = make([]int, len(texts))
	}

	for i := range dp {
		count := -1
		for j := i; j < len(texts); j++ {
			count += len(texts[j]) + 1
			dp[i][j] = calculateCost(count, w)
		}
	}

	result := make([]int, len(texts))
	lineIndexs := make([]int, len(texts))
	for i := len(texts) - 1; i >= 0; i-- {
		minCost := dp[i][len(texts)-1]
		newLineIndex := len(texts)
		if minCost == math.MaxInt16 {
			for j := len(texts) - 1; j > i; j-- {
				if dp[i][j-1] == math.MaxInt16 {
					continue
				}
				if minCost > dp[i][j-1]+result[j] {
					minCost = dp[i][j-1] + result[j]
					newLineIndex = j
				}
			}
		}
		lineIndexs[i] = newLineIndex
		result[i] = minCost
	}
	return result[0]
}
