package bst

import (
	"math"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n1, n2 := len(nums1), len(nums2)
	if n1 > n2 {
		FindMedianSortedArrays(nums2, nums1)
	}
	low, high := 0, n1
	total_lenght := len(nums1) + len(nums2)
	for low <= high {
		m1 := (low + high) / 2
		m2 := ((len(nums1) + len(nums2) + 1) / 2) - m1
		l1 := math.MinInt
		l2 := math.MinInt
		if m1 > 0 && n1 > 0 {
			l1 = nums1[m1-1]
		}
		if m2 > 0 && n2 > 0 {
			l2 = nums2[m2-1]
		}

		r1 := math.MaxInt
		r2 := math.MaxInt
		if m1 < n1 {
			r1 = nums1[m1]
		}
		if m2 < n2 {
			r2 = nums2[m2]
		}
		if l1 <= r2 && l2 <= r1 {
			if total_lenght%2 == 0 {
				return float64((max(l1, l2) + min(r1, r2)) / 2)
			} else {
				return float64(max(l1, l2))
			}
		} else if l1 > r2 {
			high = m1 - 1
		} else {
			low = m1 + 1
		}
	}
	return 0.0
}

// 7, 12, 14, 15, 16
