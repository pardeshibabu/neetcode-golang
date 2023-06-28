package scripts

func maxOfTwo(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func CharacterReplacement(s string, k int) int {
	l, r := 0, 0
	maxF := 0
	count := make(map[string]int)
	res := 0
	for r < len(s) {
		count[string(s[r])]++
		maxF = maxOfTwo(maxF, count[string(s[r])])
		// fmt.Println("maxF", maxF, count)
		for ((r - l + 1) - maxF) > k {
			count[string(s[l])]--
			l++
		}
		res = maxOfTwo(res, (r - l + 1))
		r++
	}
	return res
}
