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
	gb.Forward(5)
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

func TestGapBuffer_Empty(t *testing.T) {
	gb := NewGapBuffer("")
	if gb.GapLength == 0 {
		t.Fatalf("Zero gap length.")
	}

	gb.Forward(1)
	if gb.GapIndex != 0 {
		t.Fatalf("Forward should be a NOOP. GapIndex=%d, expected=1", gb.GapIndex)
	}

	gb.Backward(1)
	if gb.GapIndex != 0 {
		t.Fatalf("Backward should be a NOOP. GapIndex=%d, expected=0", gb.GapIndex)
	}

	gb.Delete(1)
	if gb.GapIndex != 0 {
		t.Fatalf("Delete should be a NOOP. GapIndex=%d", gb.GapIndex)
	}
}

func TestGapBuffer_SingleCharacter(t *testing.T) {
	text := "x"
	gb := NewGapBuffer(text)
	if gb.GapIndex != 1 {
		t.Fatalf("GapIndex=%d, expected=1", gb.GapIndex)
	}

	gb.Forward(1)
	if gb.GapIndex != 1 {
		t.Fatalf("Forward should be a NOOP. GapIndex=%d, expected=1", gb.GapIndex)
	}

	gb.Backward(1)
	if gb.GapIndex != 0 {
		t.Fatalf("Backward to start of buffer. GapIndex=%d, expected=0", gb.GapIndex)
	}

	gb.Backward(1)
	if gb.GapIndex != 0 {
		t.Fatalf("Backward should be a NOOP. GapIndex=%d, expected=0", gb.GapIndex)
	}

	gb.Forward(2)
	if gb.GapIndex != len(text) {
		t.Fatalf("Forward to end of text.  GapIndex=%d, expected=1", gb.GapIndex)
	}
}
