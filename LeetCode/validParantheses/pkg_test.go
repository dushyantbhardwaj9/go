package main

import "testing"

func TestIsValidParantheses(t *testing.T) {
	actualValue := IsValidParantheses("()[]{}")
	expectedValue := true
	if actualValue != expectedValue {
		t.Errorf("Expectd Value(%t) is not same as actual value(%t)", expectedValue, actualValue)
	}
}

func Test2IsValidParantheses(t *testing.T) {
	actualValue := IsValidParantheses("(]")
	expectedValue := false
	if actualValue != expectedValue {
		t.Errorf("Expectd Value(%t) is not same as actual value(%t)", expectedValue, actualValue)
	}
}
