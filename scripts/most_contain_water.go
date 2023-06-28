package scripts

import "math"

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func MaxArea(height []int) int {
	l := 0
	r := len(height) - 1
	max := math.MinInt
	for l < r {
		tempMax := (Min(height[l], height[r]) * (r - l))
		max = Max(max, tempMax)
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
