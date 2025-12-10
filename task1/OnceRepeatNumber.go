package task1

func getOnceRepeatNum(arr []int) int {
	mapping := make(map[int]int)
	for i := range arr {
		num := arr[i]
		count := mapping[num]
		count++
		mapping[num] = count
	}

	for k, v := range mapping {
		if v == 1 {
			return k
		}
	}
	return 0
}
