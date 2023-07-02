package main

import (
	"fmt"
	stackneetcode "neetcode/stack_neetcode"
)

func main() {
	fmt.Println("Hello world")
	var slice = []int{2, 1, 5, 6, 2, 3}
	// fmt.Println(scripts.LongestConsecutive(slice))
	// s1 := "()[]{}"
	// s2 := "ab"
	// var arrS = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(stackneetcode.LargestRectangleArea(slice))
	return
}

// permutation_in_string due
// 0 [] -1
// 1 [0] 0
// 2 [1] 0
// [1] 0
// 3 [1 2] 1
// [1 2] 1
// 4 [1 2 3] 2
// 5 [1 2 4] 2
// [1 2 4] 2
// [1 2 4 5]
// stack
// [1 0 0 4 0 0]
// [1 0 0 4 0 0]
