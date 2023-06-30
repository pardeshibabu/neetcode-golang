package main

import (
	"fmt"
	stackneetcode "neetcode/stack_neetcode"
)

func main() {
	fmt.Println("Hello world")
	// var slice = []int{1, 3, -1, -3, 5, 3, 6, 7}
	// fmt.Println(scripts.LongestConsecutive(slice))
	// s1 := "()[]{}"
	// s2 := "ab"
	var arrS = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(stackneetcode.EvalRPN(arrS))
	return
}

// permutation_in_tring due
