package task1

import "testing"

func TestCycleNumber(t *testing.T) {
	if !isCycleNumber(123454321) {
		t.Error("The num is a palindromic number")
	}
}

func TestNonCycleNumber(t *testing.T) {
	if isCycleNumber(12345432) {
		t.Error("The num is not a palindromic number")
	}
}
