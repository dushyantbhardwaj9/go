package main

import "testing"

func TestRomanToInt(t *testing.T) {
	actualValue := romanToInt("MCMXCIV")
	expectedValue := 1994
	if actualValue != expectedValue {
		t.Errorf("Expectd Value(%d) is not same as actual string (%d)", expectedValue, actualValue)
	}
}
