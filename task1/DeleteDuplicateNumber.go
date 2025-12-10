package task1

func deleteDuplicateNumber(arr []int) int {
	l := len(arr)
	if l <= 1 {
		// 无元素或仅有一个元素，无需删除
		return l
	}
	// 慢指针，从第一个数字开始，无重复
	slow := 0
	// 快指针，从第二个数字开始判断是否有重复
	fast := 1
	// 1. arr[fast] != arr[slow], 说明无重复，arr[slow + 1] = arr[fast], slow++, fast ++
	// 2. arr[fast] == arr[slow], 说明重复，向后遍历，fast++

	for ; fast < len(arr); fast++ {
		if arr[fast] != arr[slow] {
			slow++
			arr[slow] = arr[fast]
		}
	}

	// 将slow位置后的所有元素设置为0值
	for i := slow + 1; i < fast; i++ {
		arr[i] = 0
	}
	return slow + 1
}
