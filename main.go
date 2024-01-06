package main

import (
	"fmt"
	"neetcode/bst"
)

func main() {
	fmt.Println("Hello world testing commit")
	// var slice = []int{13, 14, 15, -1, 0, 3, 5, 9, 12}
	// // fmt.Println(scripts.LongestConsecutive(slice))
	// // s1 := "()[]{}"
	// // s2 := "ab"
	// // var arrS = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	// fmt.Println(bst.FindMin(slice))

	//  Your TimeMap object will be instantiated and called as such:
	obj := bst.Constructor()
	obj.Set("foo", "bar", 2)
	// param_2 := obj.Get("foo", 1)
	// fmt.Println("finished::", param_2)
	obj.Set("foo", "bar2", 4)
	obj.Set("foo", "bar3", 6)
	obj.Set("foo", "bar4", 8)
	// fmt.Println("f:::", obj.Get("foo", 4))
	fmt.Println(obj.Get("foo", 9))
	return
}

// permutation_in_string due
