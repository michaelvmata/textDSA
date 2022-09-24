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
	marker := position + 1
	for i, character := range text {
		if i <= marker {
			ds.before = append(ds.before, character)
		} else {
			ds.after = append(ds.after, character)
		}
	}
	return &ds
}


func (ds *DoubleStack) Text() string {
	var b strings.Builder
	b.Grow(len(ds.before) + len(ds.after))
	for _, character := range ds.before {
		b.WriteRune(character)
	}
	for _, character := range ds.after {
		b.WriteRune(character)
	}
	return b.String()
}

func (ds *DoubleStack) Insert(position int, s string) {

}

func (ds *DoubleStack) Delete(position int, value string) {

}

func (ds *DoubleStack) Skip(position int) {

}


