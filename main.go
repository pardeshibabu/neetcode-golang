package main

import (
	"container/heap"
	"fmt"
	heappriorityqueue "neetcode/heap-priority-queue"
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

	// obj1 := linkedlist.Node{
	// 	Val:    1,
	// 	Next:   nil,
	// 	Random: nil,
	// }
	// obj2 := linkedlist.Node{
	// 	Val:    2,
	// 	Next:   nil,
	// 	Random: nil,
	// }
	// obj3 := linkedlist.Node{
	// 	Val:    3,
	// 	Next:   nil,
	// 	Random: nil,
	// }
	// obj4 := linkedlist.Node{
	// 	Val:    4,
	// 	Next:   nil,
	// 	Random: nil,
	// }
	// obj5 := linkedlist.Node{
	// 	Val:    5,
	// 	Next:   nil,
	// 	Random: nil,
	// }
	// obj1.Next = &obj2
	// obj2.Next = &obj3
	// obj3.Next = &obj4
	// obj4.Next = &obj5

	// obj1.Random = nil
	// obj2.Random = &obj1
	// obj3.Random = &obj5
	// obj4.Random = &obj3
	// obj5.Random = &obj1
	// t := &obj1
	// for t != nil {
	// 	fmt.Print(t.Val, "-->")
	// 	t = t.Next
	// }
	// fmt.Println()

	// an array is not heapify yet
	arrayForInit := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}

	minHeap := &heappriorityqueue.MinHeap{}
	*minHeap = arrayForInit
	heap.Init(minHeap)
	k := 3
	for len(*minHeap) > k {
		minHeap.Pop()
	}
	fmt.Println("hepified array::", minHeap)
	fmt.Println("Kth largest::", minHeap)
	return
}

// permutation_in_string due
