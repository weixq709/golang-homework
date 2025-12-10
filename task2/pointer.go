package task2

func addNumber(num *int) {
	*num = *num + 10
}

func multiplyArrayElements(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = (*arr)[i] * 2
	}
}
