package main

import (
	"fmt"
	linkedlist "neetcode/linked_list"
)

func main() {
	fmt.Println("Hello world")
	// var slice1 = []int{1, 2}
	// var slice2 = []int{3, 4}
	// // fmt.Println(scripts.LongestConsecutive(slice))
	// // s1 := "()[]{}"
	// // s2 := "ab"
	// // var arrS = []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	// fmt.Println(bst.FindMedianSortedArrays(slice1, slice2))

	// //  Your TimeMap object will be instantiated and called as such:
	// obj := bst.Constructor()
	// obj.Set("foo", "bar", 2)
	// // param_2 := obj.Get("foo", 1)
	// // fmt.Println("finished::", param_2)
	// obj.Set("foo", "bar2", 4)
	// obj.Set("foo", "bar3", 6)
	// obj.Set("foo", "bar4", 8)
	// // fmt.Println("f:::", obj.Get("foo", 4))
	// fmt.Println(obj.Get("foo", 9))

	obj1 := linkedlist.ListNode{
		Val:  1,
		Next: nil,
	}
	obj2 := linkedlist.ListNode{
		Val:  2,
		Next: nil,
	}
	obj3 := linkedlist.ListNode{
		Val:  3,
		Next: nil,
	}
	obj4 := linkedlist.ListNode{
		Val:  4,
		Next: nil,
	}
	obj5 := linkedlist.ListNode{
		Val:  5,
		Next: nil,
	}
	obj1.Next = &obj2
	obj2.Next = &obj3
	obj3.Next = &obj4
	obj4.Next = &obj5
	linkedlist.ReorderList(&obj1)
	t := &obj1
	for t != nil {
		fmt.Print(t.Val, "-->")
		t = t.Next
	}
	return
}

// permutation_in_string due
