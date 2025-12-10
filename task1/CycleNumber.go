package task1

func isCycleNumber(num int) bool {
	arr := make([]int, 0)
	for {
		n := num % 10
		num = num / 10
		arr = append(arr, n)

		if num == 0 {
			break
		}
	}

	l := len(arr)
	for i := 0; i < l/2; i++ {
		if arr[i] != arr[l-i-1] {
			return false
		}
	}
	return true
}
