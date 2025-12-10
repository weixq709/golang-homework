package task1

import "testing"

func TestValidBracket(t *testing.T) {
	if !isValidBracket("[{()}]") {
		t.Errorf("This is a valid string")
	}
}

func TestInvalidBracket(t *testing.T) {
	if isValidBracket("{(]]") {
		t.Errorf("This is not a valid string")
	}
}
