package textDSA

import (
	"strings"
)

type DoubleStack struct {
	before []rune
	after []rune
	position int
}

func NewDoubleStack(position int, text string) *DoubleStack {
	ds := DoubleStack{
		before: make([]rune, 0),
		after: make([]rune, 0),
		position: position,
	}
	for i, character := range text {
		if i < position {
			ds.before = append(ds.before, character)
		} else {
			ds.after = append(ds.after, character)
		}
	}

	for i := 0; i < len(ds.after) / 2; i++ {
		j := len(ds.after) - i - 1
		ds.after[i], ds.after[j] = ds.after[j], ds.after[i]
	}
	return &ds
}


func (ds *DoubleStack) Text() string {
	var b strings.Builder
	b.Grow(len(ds.before) + len(ds.after))
	for _, character := range ds.before {
		b.WriteRune(character)
	}
	for k := range ds.after {
		character := ds.after[len(ds.after) - k - 1]
		b.WriteRune(character)
	}
	return b.String()
}

func (ds *DoubleStack) Insert(position int, s string) {

}

func (ds *DoubleStack) Delete(position int, value string) {

}

func (ds *DoubleStack) Skip(position int) {
	for position < ds.position {
		l := len(ds.before)
		character := ds.before[l-1]
		ds.before = ds.before[:l-1]
		ds.after = append(ds.after, character)
		ds.position -= 1
	}
	for position > ds.position {
		l := len(ds.after)
		character := ds.after[l-1]
		ds.after = ds.after[:l-1]
		ds.before = append(ds.before, character)
		ds.position += 1
	}
}