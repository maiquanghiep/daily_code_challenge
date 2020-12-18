package textjustification

import "testing"

func TestTextJustification(t *testing.T) {
	words := []string{
		"Imagination", "is", "more", "important", "than", "knowledge.",
	}
	maxWidth := 14
	expectedResult := []string{
		"Imagination is", "more important", "than          ", "knowledge.    ",
	}

	actualResult := fullJustify(words, maxWidth)
	for i, v := range actualResult {
		if v != expectedResult[i] {
			t.Errorf("Expected: %s but got: %s", expectedResult[i], v)
			return
		}
	}

	words1 := []string{
		"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country",
	}
	maxWidth1 := 16
	expectedResult1 := []string{
		"ask   not   what", "your country can", "do  for  you ask", "what  you can do", "for your country",
	}

	actualResult1 := fullJustify(words1, maxWidth1)
	for i, v := range actualResult1 {
		if v != expectedResult1[i] {
			t.Errorf("Expected: %s but got: %s", expectedResult1[i], v)
			return
		}
	}

	words2 := []string{
		"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do",
	}
	maxWidth2 := 20
	expectedResult2 := []string{
		"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  ",
	}

	actualResult2 := fullJustify(words2, maxWidth2)
	for i, v := range actualResult2 {
		if v != expectedResult2[i] {
			t.Errorf("Expected: %s but got: %s", expectedResult2[i], v)
			return
		}
	}
}
