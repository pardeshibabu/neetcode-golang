package scripts

func TwoSum(numbers []int, target int) []int {
	var a []int
	n := len(numbers)
	i, j := 0, n-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			a = append(a, i+1)
			a = append(a, j+1)
			break
		}
		if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return a
}
