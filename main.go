package main

import (
	"fmt"
	stackneetcode "neetcode/stack_neetcode"
)

func main() {
	fmt.Println("Hello world")
	var slice = []int{73, 74, 75, 71, 69, 72, 76, 73}
	// fmt.Println(scripts.LongestConsecutive(slice))
	// s1 := "()[]{}"
	// s2 := "ab"
	// var arrS = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(stackneetcode.DailyTemperatures(slice))
	return
}

// permutation_in_tring due
// Hello world
// 0 8 89 100
// 1 2 62 70
// 2 7 70 76
// 3 7 58 76
// 4 7 47 76
// 5 7 47 76
// 6 7 46 76
// 7 8 76 100
// panic: runtime error: index out of range [10] with length 10
