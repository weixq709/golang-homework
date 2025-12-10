package task1

var mapping = map[string]string{
	"]": "[",
	")": "(",
	"}": "{",
}

func isValidBracket(str string) bool {
	stack := make([]string, 0)
	for _, ch := range str {
		tag := string(ch)
		openTag, exists := mapping[tag]
		// 存在开标签且最后一个是开标签，出栈，否则入栈
		pop := exists && stack[len(stack)-1] == openTag

		if pop {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, tag)
		}
	}
	return len(stack) == 0
}
