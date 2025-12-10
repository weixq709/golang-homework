package task1

import "math"

func mergerRange(arr [][]int) [][]int {
	l := len(arr)
	if l <= 1 {
		// 无需合并
		return arr
	}
	merged := make([][]int, 0)
	// 从弟二个区间开始处理，判断start是否在前一个区间内，如果满足条件，则合并
	// 合并后的区间起始值为前一个区间的起始值，结束值为前一个区间的结束值与当前区间结束值的最大值
	// newRange: [res[i - 1][0], max(res[i - 1][1], res[i[[1])]
	// 如果不在前一个区间，则添加此区间作为合并结果，更新可合并区间为当前区间

	// 设置待合并区间为第一个区间
	merging := arr[0]

	for i := 1; i < l; i++ {
		current := arr[i]
		start, end := current[0], current[1]
		prevStart, prevEnd := merging[0], merging[1]

		if start >= prevStart && start <= prevEnd {
			// 在前一个区间内，可以合并
			merging[0] = prevStart
			merging[1] = int(math.Max(float64(prevEnd), float64(end)))
			continue
		}
		if end >= prevStart && end <= prevEnd {
			merging[0] = start
			merging[1] = int(math.Max(float64(prevEnd), float64(end)))
			continue
		}
		// 添加已合并区间到合并结果
		merged = append(merged, merging)
		merging = current
	}
	// 添加最后一次合并结果
	merged = append(merged, merging)
	return merged
}
