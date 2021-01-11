package zigzag

import (
	"strings"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"
*/

/*
There are 2 ways to convert the string:
- From left to right, put each char into the corresponding row. Use one variable to check the direction (Up or down)
- For each row, get all char belongs to that row.
I solved it using the second solution. First I devided the zigzag into 3 parts: Top, Middle, and Bottom.
Then I concat 3 parts into the result
It's actually can be solved in only 1 for loop but I find it diffcult to read.
*/
func convert(s string, numRows int) string {
	var sb strings.Builder

	if numRows == 1 {
		return s
	}

	for pos := 0; pos < len(s); pos = pos + (numRows-2)*2 + 2 {
		sb.WriteByte(s[pos])
	}

	for i := 1; i < numRows-1; i++ {
		for pos := i; pos < len(s); pos = pos + (numRows-2)*2 + 2 {
			sb.WriteByte(s[pos])
			middle := pos + (numRows-i-1)*2
			if middle < len(s) {
				sb.WriteByte(s[middle])
			}
		}
	}

	for pos := numRows - 1; pos < len(s); pos = pos + (numRows-2)*2 + 2 {
		sb.WriteByte(s[pos])
	}
	return sb.String()
}
