func isValid(s string) bool {
	stack := []rune{}
	hash := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, char := range s {
		if match, found := hash[char]; found {
			if len(stack) > 0 && stack[len(stack)-1] == match {
				stack = stack[:len(stack)-1] // Pop
			} else {
				return false
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0

}