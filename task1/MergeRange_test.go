package task1

import (
	"testing"
)

func TestMergeRange1(t *testing.T) {
	arr := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	match := [][]int{{1, 6}, {8, 10}, {15, 18}}
	res := mergerRange(arr)
	if !isMatch(res, match) {
		t.Error("mergerRange error")
	}
}

func TestMergeRange2(t *testing.T) {
	arr := [][]int{{1, 4}, {4, 5}}
	match := [][]int{{1, 5}}
	res := mergerRange(arr)
	if !isMatch(res, match) {
		t.Error("mergerRange error")
	}
}

func TestMergeRange3(t *testing.T) {
	arr := [][]int{{4, 7}, {1, 4}}
	match := [][]int{{1, 7}}
	res := mergerRange(arr)
	if !isMatch(res, match) {
		t.Error("mergerRange error")
	}
}

func isMatch(res [][]int, match [][]int) bool {
	for i := 0; i < len(match); i++ {
		for j := 0; j < len(match[i]); j++ {
			if match[i][j] != res[i][j] {
				return false
			}
		}
	}
	return true
}
