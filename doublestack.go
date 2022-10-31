package textDSA

import (
	"strings"
)

type DoubleStack struct {
	before   []rune
	after    []rune
	position int
}

func NewDoubleStack(text string) *DoubleStack {
	ds := DoubleStack{
		before:   make([]rune, 0),
		after:    []rune(text),
		position: 0,
	}
	for i := 0; i < len(ds.after)/2; i++ {
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
		character := ds.after[len(ds.after)-k-1]
		b.WriteRune(character)
	}
	return b.String()
}

func (ds *DoubleStack) Insert(s string) {
	text := []rune(s)
	for i := 0; i < len(text)/2; i++ {
		j := len(text) - i - 1
		text[i], text[j] = text[j], text[i]
	}
	for k := range text {
		character := text[len(s)-k-1]
		ds.before = append(ds.before, character)
		ds.position++
	}
}

func (ds *DoubleStack) Delete(count int) {
	for i := 0; i < count && len(ds.after) > 0; i++ {
		l := len(ds.after)
		ds.after = ds.after[:l-1]
	}
}

func (ds *DoubleStack) Skip(count int) {
	if count > 0 {
		ds.Forward(count)
	} else if count < 0 {
		ds.Backward(-count)
	}
}

func (ds *DoubleStack) Backward(count int) {
	for count > 0 && ds.position > 0 {
		l := len(ds.before)
		character := ds.before[l-1]
		ds.before = ds.before[:l-1]
		ds.after = append(ds.after, character)
		ds.position -= 1
		count -= 1
	}
}

func (ds *DoubleStack) Forward(count int) {
	for count > 0 && len(ds.after) > 0 {
		l := len(ds.after)
		character := ds.after[l-1]
		ds.after = ds.after[:l-1]
		ds.before = append(ds.before, character)
		ds.position += 1
		count -= 1
	}
}

func (ds DoubleStack) Peek() string {
	l := len(ds.after)
	if l > 0 {
		return string(ds.after[l-1])
	}
	return ""
}
