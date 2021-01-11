package zigzag

import "testing"

func TestConvertZigZag(t *testing.T) {
	s := "PAYPALISHIRING"
	numRows := 3
	r := convert(s, numRows)

	if r != "PAHNAPLSIIGYIR" {
		t.Errorf("Expected PAHNAPLSIIGYIR, but got %s", r)
	}

	numRows = 4
	r = convert(s, numRows)
	if r != "PINALSIGYAHRPI" {
		t.Errorf("Expected PAHNAPLSIIGYIR, but got %s", r)
	}

	numRows = 1
	r = convert("A", numRows)
	if r != "A" {
		t.Errorf("Expected a, but got %s", r)
	}
}
