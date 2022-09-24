package textDSA

import "testing"

func TestDoubleStack(t *testing.T) {
	text := "foobar"
	ds := NewDoubleStack(0, text)
	if ds.Text() != text {
		t.Fatalf("DoubleStack Text()=%s, expected=%s", ds.Text(), text)
	}
	if len(ds.before) != 0 {
		t.Fatalf("DoubleStack len(ds.before)=%d, expected=%d", len(ds.before), 0)
	}
	if len(ds.after) != len(text) {
		t.Fatalf("DoubleStack len(ds.after)=%d, expected=%d", len(ds.after), len(text))
	}
	ds.Insert(0, " ")
	ds.Delete(0, " ")
	ds.Skip(1)
}