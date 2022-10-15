package textDSA

import (
	"reflect"
	"testing"
)

type Operation int8

const (
	Insert Operation = iota
	Delete
	Skip
)

func (o Operation) String() string {
	switch o {
	case Insert:
		return "Insert"
	case Delete:
		return "Delete"
	case Skip:
		return "Skip"
	}
	return "Unknown"
}

type OperationalTransformation struct {
	operation Operation
	count     int
	text      string
}

type OTTestCase struct {
	label                      string
	original                   string
	expected                   string
	position                   int
	operationalTransformations []OperationalTransformation
}

func NewOTTestCase(label string, original string, expected string, position int) OTTestCase {
	return OTTestCase{
		label:                      label,
		original:                   original,
		expected:                   expected,
		position:                   position,
		operationalTransformations: make([]OperationalTransformation, 0),
	}
}

func BuildBatteryOfTestCases() []OTTestCase {
	battery := make([]OTTestCase, 0)
	battery = append(battery, buildBlankTestCase())
	battery = append(battery, buildDeleteTestCase())
	battery = append(battery, buildInsertTestCase())
	return battery
}

func buildBlankTestCase() OTTestCase {
	ott := NewOTTestCase("Blank text", "", "", 0)
	ott.operationalTransformations = append(ott.operationalTransformations, OperationalTransformation{operation: Delete, count: 5})
	ott.operationalTransformations = append(ott.operationalTransformations, OperationalTransformation{operation: Skip, count: -5})
	return ott
}

func buildDeleteTestCase() OTTestCase {
	ott := NewOTTestCase("Multi Delete", "abcd", "a", 0)
	ott.operationalTransformations = append(ott.operationalTransformations, OperationalTransformation{operation: Skip, count: 1})
	ott.operationalTransformations = append(ott.operationalTransformations, OperationalTransformation{operation: Delete, count: 2})
	ott.operationalTransformations = append(ott.operationalTransformations, OperationalTransformation{operation: Delete, count: 1})
	return ott
}

func buildInsertTestCase() OTTestCase {
	ott := NewOTTestCase("Multi insert", "a", "dcba", 0)
	ott.operationalTransformations = append(ott.operationalTransformations, OperationalTransformation{operation: Insert, text: "b"})
	ott.operationalTransformations = append(ott.operationalTransformations, OperationalTransformation{operation: Insert, text: "dc"})
	return ott
}

type Transformer interface {
	Insert(string)
	Delete(int)
	Skip(int)
	Text() string
}

func TestBattery(t *testing.T) {
	RunBattery(t, func(position int, original string) Transformer { return NewDoubleStack(position, original) })
}

func RunBattery(t *testing.T, newTransformer func(int, string) Transformer) {
	battery := BuildBatteryOfTestCases()
	for _, testCase := range battery {
		transformer := newTransformer(testCase.position, testCase.original)
		for _, ot := range testCase.operationalTransformations {
			switch ot.operation {
			case Insert:
				transformer.Insert(ot.text)
			case Delete:
				transformer.Delete(ot.count)
			case Skip:
				transformer.Skip(ot.count)
			}
		}
		if transformer.Text() != testCase.expected {
			t.Fatalf("Transformer %s TestCase label=%s, ds.Text()=%s, expected=%s", reflect.TypeOf(transformer).Elem().Name(), testCase.label, transformer.Text(), testCase.expected)
		}
	}
}
