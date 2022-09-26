package textDSA

import "testing"

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
	count int
	text string
}

type OTTestCase struct {
	label string
	original string
	expected string
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
	return battery
}

func buildBlankTestCase() OTTestCase {
	ott := NewOTTestCase("Blank text", "", "", 0)
	ott.operationalTransformations = append(ott.operationalTransformations, OperationalTransformation{operation: Delete, count: 5})
	ott.operationalTransformations = append(ott.operationalTransformations, OperationalTransformation{operation: Skip, count: -5})
	return ott
}

func TestDoubleStackBattery(t *testing.T) {
	battery := BuildBatteryOfTestCases()
	for _, testCase := range battery {
		ds := NewDoubleStack(testCase.position, testCase.original)
		for _, ot := range testCase.operationalTransformations {
			switch ot.operation {
			case Insert:
				ds.Insert(ot.text)
			case Delete:
				ds.Delete(ot.count)
			case Skip:
				ds.Skip(ot.count)
			}
		}
		if ds.Text() != testCase.expected {
			t.Fatalf("TestCase label=%s, ds.Text()=%s, expected=%s", testCase.label, ds.Text(), testCase.expected)
		}
	}
}