package utils

import (
	"fmt"
	"testing"
)

func TestColorHandlesStringInput(t *testing.T) {
	input := "test"
	expectedRed := fmt.Sprintf("%s%s%s", ColorRed, input, ColorDefault)
	resultRed := Red(input)

	if resultRed != expectedRed {
		t.Errorf("Expected %s, but got %s", expectedRed, resultRed)
	}

	expectedGreen := fmt.Sprintf("%s%s%s", ColorGreen, input, ColorDefault)
	resultGreen := Green(input)

	if expectedGreen != resultGreen {
		t.Errorf("Expected %s, but got %s", expectedGreen, resultGreen)
	}

	expectedBlue := fmt.Sprintf("%s%s%s", ColorBlue, input, ColorDefault)
	resultBlue := Blue(input)

	if expectedBlue != resultBlue {
		t.Errorf("Expected %s, but got %s", expectedBlue, resultBlue)
	}

	expectedGray := fmt.Sprintf("%s%s%s", ColorGray, input, ColorDefault)
	resultGray := Gray(input)

	if expectedGray != resultGray {
		t.Errorf("Expected %s, but got %s", expectedGray, resultGray)
	}
}
