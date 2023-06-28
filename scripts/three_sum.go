package scripts

func merg(arr []int, l, m, r int) {
	nl := (m - l + 1)
	rl := (r - m)
	left := make([]int, nl)
	right := make([]int, rl)
	for i := 0; i < nl; i++ {
		left[i] = arr[l+i]
	}
	for j := 0; j < rl; j++ {
		right[j] = arr[m+1+j]
	}
	i, j, k := 0, 0, l
	for i < nl && j < rl {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < nl {
		arr[k] = left[i]
		i++
		k++
	}
	for j < rl {
		arr[k] = right[j]
		j++
		k++
	}
}

func MergeSort(arr []int, l, r int) {
	if l < r {
		mid := l + (r-l)/2
		MergeSort(arr, l, mid)
		MergeSort(arr, mid+1, r)
		merg(arr, l, mid, r)
	}
}

func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	MergeSort(nums, 0, len(nums)-1)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := n - 1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				var temp = []int{}
				temp = append(temp, nums[i], nums[l], nums[r])
				result = append(result, temp)
				l++
				for l < r && nums[l-1] == nums[l] {
					l++
				}
			} else if nums[i]+nums[l]+nums[r] > 0 {
				r--
			} else {
				l++
			}
		}
	}
	return result
}
