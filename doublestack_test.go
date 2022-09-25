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

	skip := 2
	ds.Skip(skip)
	if ds.Text() != text {
		t.Fatalf("DoubleStack Text()=%s, expected=%s", ds.Text(), text)
	}
	if len(ds.before) != skip {
		t.Fatalf("DoubleStack len(ds.before)=%d, expected=%d", len(ds.before), 2)
	}
	if len(ds.after) != len(text) - skip {
		t.Fatalf("DoubleStack len(ds.after)=%d, expected=%d", len(ds.after), len(text)-skip)
	}

	skip = len(text) - skip
	ds.Skip(skip)
	if ds.Text() != text {
		t.Fatalf("DoubleStack Text()=%s, expected=%s", ds.Text(), text)
	}
	if len(ds.before) != len(text) {
		t.Fatalf("DoubleStack len(ds.before)=%d, expected=%d", len(ds.before), len(text))
	}
	if len(ds.after) != 0 {
		t.Fatalf("DoubleStack len(ds.after)=%d, expected=%d", len(ds.after), 0)
	}


	space := " "
	ds.Insert(space)
	if ds.Text() != text+space {
		t.Fatalf("DoubleStack Text()=%s, expected=%s", ds.Text(), text+space)
	}

	ds.Delete(1)
	if ds.Text() != text {
		t.Fatalf("DoubleStack Text()=%s, expected=%s", ds.Text(), text)
	}
}