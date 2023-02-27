package main

import (
	"testing"
)

func testGPA(t *testing.T) {
	testInput := args{
		"english",
		"cp",
		5,
		88,
	}

	credits, points := testInput.calcGPA()
	if gpa := credits / points; credits != 5 || points != 16.5 || gpa != 3.3 {
		t.Error("Unexpected output")
	}
}
