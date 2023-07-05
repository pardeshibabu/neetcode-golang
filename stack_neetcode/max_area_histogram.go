package stackneetcode

import (
	"math"
)

func NextSmallestRightIndex(arr []int, n int) []int {
	res := make([]int, n)
	var stack []int
	r := 1
	stack = append(stack, 0)
	for r < n {
		for len(stack) > 0 && arr[stack[len(stack)-1]] > arr[r] {
			top := stack[len(stack)-1]
			res[top] = r
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, r)
		r++
	}
	for i := 0; i < len(stack); i++ {
		res[stack[i]] = n
	}
	return res
}

func NextSmallestLeftIndex(a []int, n int) []int {
	res := make([]int, n)
	var stack []int
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && a[i] < a[stack[len(stack)-1]] {
			res[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)

	}
	for j := 0; j < len(stack); j++ {
		res[stack[j]] = -1
	}
	return res
}
func LargestRectangleArea(heights []int) int {
	n := len(heights)
	nsL := NextSmallestLeftIndex(heights, n)
	nsR := NextSmallestRightIndex(heights, n)
	maxArea := math.MinInt
	for i := 0; i < n; i++ {
		tempMax := ((nsR[i] - nsL[i]) - 1) * heights[i]
		if tempMax > maxArea {
			maxArea = tempMax
		}
	}
	return maxArea
}
