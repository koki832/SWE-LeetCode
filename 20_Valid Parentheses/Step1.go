func isValid(s string) bool {

	stack := ""
	for _, r := range s {
		if r == ')' {
			//　stackの最後が(の場合、Stackの最後の文字を削除
			//　それ以外はfalse
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if r == '}' {
			//　stackの最後が{の場合、Stackの最後の文字を削除
			//　それ以外はfalse
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if r == ']' {
			//　stackの最後が[の場合、Stackの最後の文字を削除
			//　それ以外はfalse
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack += string(r)
		}
	}
    if len(stack) == 0 {
        return true
    }
	return false
}