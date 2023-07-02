package stackneetcode

import (
	"fmt"
	"math"
)

func NextSmallestRightIndex(a []int, n int) []int {
	res := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		fmt.Println(i, stack, len(stack)-1)
		if len(stack) == 0 {
			stack = append(stack, i)
		} else if a[i] >= a[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && a[i] < a[len(stack)-1] {
				res[stack[len(stack)-1]] = i
				fmt.Println("before::", stack, len(stack)-1)
				stack = stack[:len(stack)-1]
				fmt.Println("after::", stack, a[len(stack)-1], a[i])
			}
			stack = append(stack, i)
		}
	}
	fmt.Println(stack)
	fmt.Println("stack")
	for j := 0; j > len(stack); j++ {
		res[stack[j]] = n
	}
	return res
}

func NextSmallestLeftIndex(a []int, n int) []int {
	res := make([]int, n)
	var stack []int
	for i := 0; i < n; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else if a[i] >= a[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && a[i] < a[len(stack)-1] {
				res[stack[len(stack)-1]] = i
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	for j := 0; j > len(stack); j++ {
		res[stack[j]] = -1
	}
	return res
}
func LargestRectangleArea(heights []int) int {
	n := len(heights)
	nsL := NextSmallestLeftIndex(heights, n)
	nsR := NextSmallestRightIndex(heights, n)
	fmt.Println(nsL)
	fmt.Println(nsR)
	maxArea := math.MinInt
	for i := 0; i < n; i++ {
		tempMax := ((nsR[i] - nsL[i]) - 1) * heights[i]
		if tempMax > maxArea {
			maxArea = tempMax
		}
	}
	return maxArea
}
