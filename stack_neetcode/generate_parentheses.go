package stackneetcode

func recursiveF(openP, closeP int, output string, stack *[]string) {
	if openP == 0 && closeP == 0 {
		*stack = append(*stack, output)
		return
	}

	if openP > 0 {
		recursiveF(openP-1, closeP, output+"(", stack)
	}

	if closeP > openP {
		recursiveF(openP, closeP-1, output+")", stack)
	}
}

func GenerateParenthesis(n int) []string {
	openP := n
	closeP := n
	var stack []string
	output := ""
	recursiveF(openP, closeP, output, &stack)
	return stack
}
