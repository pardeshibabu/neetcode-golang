package stackneetcode

import "strconv"

func EvalRPN(tokens []string) int {
	var stack []int
	for i := 0; i < len(tokens); i++ {
		if string(tokens[i]) == "+" {
			top1 := stack[len(stack)-1]
			top2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, (top1 + top2))
		} else if string(tokens[i]) == "-" {
			top1 := stack[len(stack)-1]
			top2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, (top2 - top1))
		} else if string(tokens[i]) == "*" {
			top1 := stack[len(stack)-1]
			top2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, (top2 * top1))
		} else if string(tokens[i]) == "/" {
			top1 := stack[len(stack)-1]
			top2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, (top2 / top1))
		} else {
			v, _ := strconv.Atoi(tokens[i])
			stack = append(stack, v)
		}
	}
	return stack[0]
}
