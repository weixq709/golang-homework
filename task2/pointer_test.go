package task2

import (
	"testing"
)

func TestAddNumber(t *testing.T) {
	var num = 10
	addNumber(&num)
	if num != 20 {
		t.Error("addNumber error")
	}
}

func TestMultiplyArrayElements(t *testing.T) {
	var arr = []int{1, 2, 3, 4, 5}
	res := make([]int, 0, len(arr))
	res = append(res, arr...)
	multiplyArrayElements(&res)

	for i := 0; i < len(arr); i++ {
		if res[i] != arr[i]*2 {
			t.Error("multiplyArrayElements error")
			break
		}
	}
}
