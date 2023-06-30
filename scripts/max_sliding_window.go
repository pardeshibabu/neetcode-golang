package scripts

func MaxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	dequeue := []int{}
	l, r := 0, 0
	for r < len(nums) {
		for len(dequeue) > 0 && nums[dequeue[len(dequeue)-1]] < nums[r] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, r)
		if l > dequeue[0] {
			dequeue = dequeue[1:]
		}
		if (r + 1) >= k {
			res = append(res, nums[dequeue[0]])
			l++
		}
		r++
	}
	return res
}
