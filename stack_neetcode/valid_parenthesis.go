package stackneetcode

func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var stack []string
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if len(stack) == 0 {
			stack = append(stack, c)
		} else if (c == ")" && stack[len(stack)-1] == "(") ||
			(c == "}" && stack[len(stack)-1] == "{") ||
			(c == "]" && stack[len(stack)-1] == "[") {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
