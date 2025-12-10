package task1

import "testing"

func TestPlusOne(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	plueOne(arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			t.Errorf("getLongestCommonPrefix fail")
			break
		}
	}
}
