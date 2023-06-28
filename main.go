package main

import (
	"fmt"
	"neetcode/scripts"
)

func main() {
	fmt.Println("Hello world")
	// var slice = []int{2, 1, 2, 1, 0, 1, 2}
	// fmt.Println(scripts.LongestConsecutive(slice))
	s1 := "cabwefgewcwaefgcf"
	s2 := "cae"
	fmt.Println(scripts.MinWindow(s1, s2))
	return
}

// permutation_in_tring due
