package adder

import (
	"testing"
)

func Test_addNumbers(t *testing.T) {
	result := addNumbers(2, 4)
	if result != 6 {
		t.Error("incorrect result: expected 5, got", result)
	}
}
