package scripts

func Trap(height []int) int {
	l, r, leftMax, rightMax := 0, len(height)-1, 0, 0
	res := 0
	for l < r {
		if height[l] <= height[r] {
			if height[l] >= leftMax {
				leftMax = height[l]
			} else {
				res += (leftMax - height[l])
			}
			l++
		} else {
			if height[r] >= rightMax {
				rightMax = height[r]
			} else {
				res += (rightMax - height[r])
			}
			r--
		}
	}

	return res
}
