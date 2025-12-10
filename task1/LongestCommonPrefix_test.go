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
