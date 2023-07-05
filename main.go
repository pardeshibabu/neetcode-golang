package main

import (
	"fmt"
	"neetcode/bst"
)

func main() {
	fmt.Println("Hello world")
	var slice = []int{-1, 0, 3, 5, 9, 12}
	// fmt.Println(scripts.LongestConsecutive(slice))
	// s1 := "()[]{}"
	// s2 := "ab"
	// var arrS = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(bst.Search(slice, 9))
	return
}

// permutation_in_string due
