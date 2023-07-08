package bst

func SearchMatrix(mat [][]int, target int) bool {
	lm, rm, ln, rn := 0, len(mat[0])-1, 0, len(mat)-1
	for ln <= rn {
		mid := ln + (rn-ln)/2
		if mat[mid][0] <= target && mat[mid][len(mat[0])-1] >= target {
			// column traversing
			for lm <= rm {
				mid_m := lm + (rm-lm)/2
				if mat[mid][mid_m] == target {
					return true
				} else if mat[mid][mid_m] > target {
					rm = mid_m - 1
				} else {
					lm = mid_m + 1
				}
			}
			return false

		} else if mat[mid][0] > target {
			rn = mid - 1
		} else {
			ln = mid + 1
		}
	}
	return false
}
