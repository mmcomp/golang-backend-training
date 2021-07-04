package main

import "testing"

// • CAD-$1,234.56
// • CAD$-1,234.56
// • $1234.56
// • $1,234.56
// • CAD $1234.56
// • CAD $1,234.56
// • CAD$1234.56
// • CAD$1,234.56
// • $0.09
// • $.09
// • -$0.09
// • -$.09
// • $-0.09
// • $-.09
// • CAD $0.09
// • CAD $.09
// • CAD -$0.09
// • CAD -$.09
// • CAD $-0.09
// • CAD $-.09
// • CAD$0.09
// • CAD$.09
// • CAD-$0.09
// • CAD-$.09
// • CAD$-0.09
// • CAD$-.09
// • 9¢
// • -9¢
// • 123456¢
// • -123456¢
func TestCAD_ParseCAD(t *testing.T) {
	tests := []struct{
		Input string
		Output CAD
		Error error
	}{
		{
			Input: "-$1234.56",
			Output: CAD{
				cents: -123456,
			},
			Error: nil,
		},
		{
			Input: "$-1234.56",
			Output: CAD{
				cents: -123456,
			},
			Error: nil,
		},
		{
			Input: "-$1,234.56",
			Output: CAD{
				cents: -123456,
			},
			Error: nil,
		},
		{
			Input: "$-1,234.56",
			Output: CAD{
				cents: -123456,
			},
			Error: nil,
		},
		{
			Input: "CAD -$1234.56",
			Output: CAD{
				cents: -123456,
			},
			Error: nil,
		},
		{
			Input: "CAD $-1234.56",
			Output: CAD{
				cents: -123456,
			},
			Error: nil,
		},
	}

	for testNumber, test := range tests {
		cad, err := ParseCAD(test.Input)
		if cad.cents != test.Output.cents {
			t.Errorf("Test %d :  cents is %d  but was expecting %d", testNumber, cad.cents, test.Output.cents)
		}
		if err != test.Error {
			t.Errorf("Test %d :  error is %s  but was expecting %s", testNumber, err, test.Error)
		}
	}
}