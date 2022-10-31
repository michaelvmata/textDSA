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
	return &GapBuffer{
		Buffer:    []rune(text),
		GapIndex:  0,
		GapLength: 0,
	}
}

func (gb GapBuffer) Text() string {
	var b strings.Builder
	b.Grow(len(gb.Buffer) - gb.GapLength)
	for i, character := range gb.Buffer {
		if i < gb.GapIndex || i > gb.GapIndex+gb.GapLength-1 {
			b.WriteRune(character)
		}
	}
	return b.String()
}

func (gb *GapBuffer) Skip(count int) {
	if count > 0 {
		gb.Forward(count)
	} else if count < 0 {
		gb.Backward(-count)
	}
}

func (gb *GapBuffer) Forward(count int) {
	for count > 0 && gb.GapIndex+gb.GapLength < len(gb.Buffer) {
		if gb.GapLength > 0 {
			gb.Buffer[gb.GapIndex] = gb.Buffer[gb.GapIndex+gb.GapLength]
			gb.Buffer[gb.GapIndex+gb.GapLength] = 0
		}
		count--
		gb.GapIndex++
	}
}

func (gb *GapBuffer) Backward(count int) {
	for count > 0 && gb.GapIndex > 0 {
		if gb.GapLength > 0 {
			gb.Buffer[gb.GapIndex+gb.GapLength-1] = gb.Buffer[gb.GapIndex-1]
			gb.Buffer[gb.GapIndex-1] = 0
		}
		count--
		gb.GapIndex--
	}
}

func (gb *GapBuffer) GrowGap() {
	grow := len(gb.Buffer)
	if grow == 0 {
		grow = 1
	}
	buffer := make([]rune, len(gb.Buffer)+grow)
	if gb.GapIndex != 0 {
		copy(buffer, gb.Buffer[:gb.GapIndex])
	}
	if gb.GapIndex < len(gb.Buffer) {
		copy(buffer[gb.GapIndex+gb.GapLength+grow:], gb.Buffer[gb.GapIndex+gb.GapLength:])
	}
	gb.Buffer = buffer
	gb.GapLength += grow
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
	if len(gb.Buffer) == 0 || gb.GapIndex+gb.GapLength == len(gb.Buffer) {
		// No content after gap
		return
	}
	for count > 0 && gb.GapIndex+gb.GapLength < len(gb.Buffer) {
		gb.Buffer[gb.GapIndex+gb.GapLength] = 0
		gb.GapLength++
		count--
	}
}

func (gb GapBuffer) Peek() string {
	if gb.GapIndex+gb.GapLength == len(gb.Buffer) {
		return ""
	}
	return string(gb.Buffer[gb.GapIndex+gb.GapLength])
}
