package bst

func FindMin(nums []int) int {
	l, r := 0, len(nums)-1
	res := nums[0]

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= nums[0] {
			l = mid + 1
		} else {
			if nums[mid] < res {
				res = nums[mid]
			}
			r = mid - 1
		}
	}
	return res
}
