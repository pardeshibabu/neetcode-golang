package scripts

// ArrayIn ...
// function will check given element prsent in given array
func ArrayIn(arr []int, value int) bool {
	for _, a := range arr {
		if a == value {
			return true
		}
	}
	return false
}
func MaxOfTwo(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := make(map[int]bool)
	longest := 0
	for _, v := range nums {
		set[v] = true
	}
	for _, v := range nums {
		if _, ok := set[v-1]; !ok {
			length := 0
			ok1 := true
			for ok1 {
				if _, ok1 := set[v+length]; ok1 {
					length += 1
					continue
				}
				ok1 = false
			}
			longest = MaxOfTwo(length, longest)
		}
	}
	return longest
}
