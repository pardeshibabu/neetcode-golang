package bst

func BSearchT(a []int, l, r, n int) int {
	if r >= l {
		mid := l + (r-l)/2
		if a[mid] == n {
			return mid
		}
		if a[mid] < n {
			return BSearchT(a, mid+1, r, n)
		} else {
			return BSearchT(a, l, mid-1, n)
		}
	}
	return -1
}

func Search(nums []int, target int) int {
	return BSearchT(nums, 0, len(nums)-1, target)
}
