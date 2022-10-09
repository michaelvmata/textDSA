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
