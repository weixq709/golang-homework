package task1

import "testing"

func TestSumTwoNumbers(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	n1, n2 := sumTwoNumbers(arr, 9)
	if !((n1 == 2 && n2 == 7) || (n1 == 7 && n2 == 2)) {
		t.Error("sumTwoNumbers error")
	}
}
