package gocipractice

import (
	"testing"
)

func testEvenOdd(t *testing.T) {
	even := evenOrOdd(10)

	if even != "Even" {
		t.Errorf("Expected Even, got %s", even)
	}

	odd := evenOrOdd(11)

	if odd != "Odd" {
		t.Errorf("Expected Odd, got %s", even)
	}
}
