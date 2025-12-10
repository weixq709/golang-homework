package task1

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	arr := []string{"flower", "flow", "flight"}
	res := getLongestCommonPrefix(arr)
	if res != "" && res != "fl" {
		// 此数组有公共前缀
		t.Errorf("getLongestCommonPrefix fail")
	}
}

func TestNoneCommonPrefix(t *testing.T) {
	arr := []string{"dog", "racecar", "car"}
	res := getLongestCommonPrefix(arr)
	if res != "" {
		t.Errorf("This string doesn't have common prefix")
	}
}
