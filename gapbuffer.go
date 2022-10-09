package textDSA

import (
	"strings"
)

type GapBuffer struct {
	GapIndex  int
	GapLength int
	Buffer    []rune
}

func NewGapBuffer(text string) *GapBuffer {
	buffer := make([]rune, len(text)*2)
	copy(buffer, []rune(text))
	return &GapBuffer{
		Buffer:    buffer,
		GapIndex:  len(text),
		GapLength: len(text),
	}
}

func (gb GapBuffer) Text() string {
	var b strings.Builder
	b.Grow(cap(gb.Buffer) - gb.GapLength)
	for i, character := range gb.Buffer {
		if i < gb.GapIndex || i > gb.GapIndex+gb.GapLength-1 {
			b.WriteRune(character)
		}
	}
	return b.String()
}

func (gb *GapBuffer) Forward(count int) {
	for count > 0 && gb.GapIndex+gb.GapLength < cap(gb.Buffer) {
		gb.Buffer[gb.GapIndex] = gb.Buffer[gb.GapIndex+gb.GapLength]
		gb.Buffer[gb.GapIndex+gb.GapLength] = 0
		count--
		gb.GapIndex++
	}
}

func (gb *GapBuffer) Backward(count int) {
	for count > 0 && gb.GapIndex > 0 {
		gb.Buffer[gb.GapIndex+gb.GapLength-1] = gb.Buffer[gb.GapIndex-1]
		count--
		gb.GapIndex--
		gb.Buffer[gb.GapIndex] = 0
	}
}

func (gb *GapBuffer) GrowGap() {
	buffer := make([]rune, len(gb.Buffer)*2)
	copy(buffer, gb.Buffer[:gb.GapIndex])
	copy(buffer, gb.Buffer[gb.GapIndex+gb.GapLength:])
	gb.Buffer = buffer
}

func (gb *GapBuffer) Insert(s string) {
	for _, character := range s {
		if gb.GapLength == 0 {
			gb.GrowGap()
		}
		gb.Buffer[gb.GapIndex] = character
		gb.GapIndex++
		gb.GapLength--
	}
}

func (gb *GapBuffer) Delete(count int) {
	for count > 0 && gb.GapIndex+gb.GapLength-1 < cap(gb.Buffer) {
		gb.Buffer[gb.GapIndex+gb.GapLength-1] = 0
		gb.GapLength++
		count--
	}
}

func (gb GapBuffer) Peek() string {
	if gb.GapIndex+gb.GapLength == cap(gb.Buffer) {
		return ""
	}
	return string(gb.Buffer[gb.GapIndex+gb.GapLength])
}
