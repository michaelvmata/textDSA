package textDSA

import (
	"testing"
)

func TestGapBuffer(t *testing.T) {
	text := "hello, world"
	gb := NewGapBuffer(text)
	if gb.Text() != text {
		t.Fatalf("GapBuffer Text()=%s expected=%s", gb.Text(), text)
	}
	gb.Forward(len(text))
	if gb.Text() != text {
		t.Fatalf("GapBuffer Text()=%s expected=%s", gb.Text(), text)
	}
	if gb.GapIndex != len(text) {
		t.Fatalf("GapBuffer GapIndex()=%d expected=%d", gb.GapIndex, len(text))
	}
	gb.Backward(len(text))
	if gb.Text() != text {
		t.Fatalf("GapBuffer Text()=%s expected=%s", gb.Text(), text)
	}
	if gb.GapIndex != 0 {
		t.Fatalf("GapBuffer GapIndex=%d expected=0", gb.GapIndex)
	}
	prefix := "ok."
	gb.Insert(prefix)
	if gb.Text() != prefix+text {
		t.Fatalf("GapBuffer Text()=%s expected=%s", gb.Text(), prefix+text)
	}
	gb.Backward(len(prefix))
	gb.Delete(len(prefix))
	if gb.Text() != text {
		t.Fatalf("GapBuffer Text()=%s expected=%s", gb.Text(), text)
	}
}

func TestGapBuffer_Insert(t *testing.T) {
	testCases := []struct {
		original []rune
		expected []rune
		index    int
		length   int
		text     string
	}{
		{
			[]rune{},
			[]rune{65},
			0,
			0,
			"A",
		},
		{
			[]rune{66},
			[]rune{65, 66},
			0,
			0,
			"A",
		},
		{
			[]rune{65},
			[]rune{65, 66},
			1,
			0,
			"B",
		},
		{
			[]rune{0, 66},
			[]rune{65, 66},
			0,
			1,
			"A",
		},
		{
			[]rune{65, 0},
			[]rune{65, 66},
			1,
			1,
			"B",
		},
		{
			[]rune{65, 0, 0, 65},
			[]rune{65, 66, 67, 65},
			1,
			2,
			"BC",
		},
		{
			[]rune{65, 0, 0, 65},
			[]rune{65, 66, 67, 68, 0, 0, 0, 65},
			1,
			2,
			"BCD",
		},
		{
			[]rune{65, 66},
			[]rune{67, 68, 65, 66},
			0,
			0,
			"CD",
		},
	}
	for _, testCase := range testCases {
		gb := NewGapBuffer("")
		gb.Buffer = make([]rune, len(testCase.original))
		copy(gb.Buffer, testCase.original)
		for i := range gb.Buffer {
			if gb.Buffer[i] != testCase.original[i] {
				t.Fatalf("Test pre grow gap mismatch actual=%v, original=%v", gb.Buffer, testCase.original)
			}
		}
		gb.GapIndex = testCase.index
		gb.GapLength = testCase.length
		gb.Insert(testCase.text)
		if len(gb.Buffer) != len(testCase.expected) {
			t.Fatalf("Buffer length mismatch actual=%v, expected=%v", gb.Buffer, testCase.expected)
		}
		for i := range gb.Buffer {
			if gb.Buffer[i] != testCase.expected[i] {
				t.Fatalf("Test grow gap mismatch actual=%v, expected=%v", gb.Buffer, testCase.expected)
			}
		}
	}
}

func TestGapBuffer_GrowGap(t *testing.T) {
	testCases := []struct {
		original []rune
		expected []rune
		index    int
		length   int
	}{
		{
			[]rune{},
			[]rune{0},
			0,
			0,
		},
		{
			[]rune{65},
			[]rune{0, 65},
			0,
			0,
		},
		{
			[]rune{65},
			[]rune{65, 0},
			1,
			0,
		},
		{[]rune{65, 66},
			[]rune{0, 0, 65, 66},
			0,
			0,
		},
		{[]rune{65, 66},
			[]rune{65, 0, 0, 66},
			1,
			0,
		},
		{[]rune{65, 66},
			[]rune{65, 66, 0, 0},
			2,
			0,
		},
		{[]rune{65, 65, 66, 66},
			[]rune{65, 65, 00, 00, 00, 00, 66, 66},
			2,
			0,
		},
	}

	for _, testCase := range testCases {
		gb := NewGapBuffer("")
		gb.Buffer = make([]rune, len(testCase.original))
		copy(gb.Buffer, testCase.original)
		for i := range gb.Buffer {
			if gb.Buffer[i] != testCase.original[i] {
				t.Fatalf("Test pre grow gap mismatch actual=%v, original=%v", gb.Buffer, testCase.original)
			}
		}
		gb.GapIndex = testCase.index
		gb.GapLength = testCase.length
		gb.GrowGap()
		for i := range gb.Buffer {
			if gb.Buffer[i] != testCase.expected[i] {
				t.Fatalf("Test grow gap mismatch actual=%v, expected=%v", gb.Buffer, testCase.expected)
			}
		}
	}
}
