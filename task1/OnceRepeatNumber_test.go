package task1

import "testing"

func TestGetOnceRepeatNumber(t *testing.T) {
	arr := []int{1, 1, 2, 3, 3}
	res := getOnceRepeatNum(arr)
	if res != 2 {
		t.Errorf("The result is doesn't match, expect: 2, actual: %d", res)
	}
}
