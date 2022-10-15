package textDSA

import "testing"

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
