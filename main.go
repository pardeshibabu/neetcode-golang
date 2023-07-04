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
