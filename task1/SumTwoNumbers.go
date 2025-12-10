package task1

func sumTwoNumbers(arr []int, target int) (int, int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return arr[i], arr[j]
			}
		}
	}
	return 0, 0
}
