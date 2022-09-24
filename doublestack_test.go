package textDSA

import "testing"

func TestDoubleStack(t *testing.T) {
	text := "foobar"
	ds := NewDoubleStack(0, text)
	if ds.Text() != text {
		t.Fatalf("DoubleStack Text()=%s, expected=%s", ds.Text(), text)
	}
	ds.Insert(0, " ")
	ds.Delete(0, " ")
	ds.Skip(1)
}