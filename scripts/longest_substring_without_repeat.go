package scripts

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LengthOfLongestSubstring(s string) int {
	l := 0
	res := 0
	m := make(map[string]bool)
	for r := 0; r < len(s); r++ {
		for m[string(s[r])] {
			delete(m, string(s[l]))
			l++
		}
		m[string(s[r])] = true
		res = max(res, (r - l + 1))
	}
	return res
}
