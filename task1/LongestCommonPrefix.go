package task1

import "math"

func getLongestCommonPrefix(arr []string) string {
	if len(arr) == 0 {
		// 数组为空，公共前缀为空
		return ""
	}

	// 两个为一组，查找公共前缀，如果某轮结果为空，说明公共前缀为空，无需向后查找
	// 否则，使用此轮公共前缀结果，继续与后一个字符串查找公共前缀
	prefix := []rune(arr[0])
	for i := 1; i < len(arr); i++ {
		s := []rune(arr[i])
		minLen := math.Min(float64(len(prefix)), float64(len(s)))

		// 记录当前公共前缀
		res := make([]rune, 0, int(minLen))

		for j := 0; j < int(minLen); j++ {
			if prefix[j] == s[j] {
				// 字符相同，添加
				res = append(res, prefix[j])
			} else {
				// 字符不同，公共字符已查找完成
				break
			}
		}

		// 无公共前缀
		if len(res) == 0 {
			return ""
		}
		// 继续向后查找
		prefix = res
	}
	return ""
}
