package textjustification

import (
	"math"
	"strings"
)

/*
Given an array of words and a width maxWidth, format the text such that each line has exactly maxWidth characters and is fully (left and right) justified.

You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly maxWidth characters.

Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.

For the last line of text, it should be left justified and no extra space is inserted between words.

Note:

A word is defined as a character sequence consisting of non-space characters only.
Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
The input array words contains at least one word.
*/

/*
Each line's len is included: Words len and space between words len.
- First I add words into line if it the total word len fits with the maxWidth.
Another condition is there is at least n-1 space between words.
- After added words into lines, I put the space in words. Messed up with the fullyJustified by calculating the amount of spaces for each word.
It can be solved easier by putting one space per word.
*/
func fullyJustified(words []string, wordLength int, spaceLength int) string {
	if len(words) == 1 {
		return leftJustified(words, wordLength, spaceLength)
	}
	spacePerWord := int(math.Ceil(float64(spaceLength) / float64((len(words) - 1))))
	var sb strings.Builder
	for i, w := range words {
		sb.WriteString(w)
		if i < len(words)-1 {
			for j := 0; j < spacePerWord; j++ {
				sb.WriteString(" ")
			}
			spaceLength -= spacePerWord
			spacePerWord = int(math.Ceil(float64(spaceLength) / float64((len(words) - i - 2))))
		}
	}

	return sb.String()
}
func leftJustified(words []string, wordLength int, spaceLength int) string {
	var sb strings.Builder
	sb.WriteString(strings.Join(words, " "))

	spareSpace := spaceLength - len(words) + 1
	for i := 0; i < spareSpace; i++ {
		sb.WriteString(" ")
	}
	return sb.String()
}

func fullJustify(words []string, maxWidth int) []string {
	curLine := []string{}
	wordLength := 0
	result := []string{}
	for _, v := range words {
		wordLength = wordLength + len(v)
		minSpace := len(curLine)
		if len(curLine) == 0 {
			minSpace = 0
		}

		if wordLength+minSpace > maxWidth {
			wordLength = wordLength - len(v)
			result = append(result, fullyJustified(curLine, wordLength, maxWidth-wordLength))
			curLine = nil
			wordLength = len(v)
		}
		curLine = append(curLine, v)
	}
	result = append(result, leftJustified(curLine, wordLength, maxWidth-wordLength))

	return result
}
