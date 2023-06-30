package main

import (
	"fmt"
	"neetcode/scripts"
)

func main() {
	fmt.Println("Hello world")
	var slice = []int{1, 3, -1, -3, 5, 3, 6, 7}
	// fmt.Println(scripts.LongestConsecutive(slice))
	// s1 := "aabb"
	// s2 := "ab"
	fmt.Println(scripts.MaxSlidingWindow(slice, 3))
	return
}

// permutation_in_tring due
