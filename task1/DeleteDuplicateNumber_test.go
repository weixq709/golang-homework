package task1

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicateNumber(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 5, 6, 6}
	newLen := deleteDuplicateNumber(arr)
	fmt.Println("newLen:", newLen)
	fmt.Println("arr:", arr)
	if newLen != 6 {
		t.Errorf("The result is not correct, expect: %d, actual: %d", 6, newLen)
	}
}
