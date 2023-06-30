package scripts

// optimized solutions
func MinWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	tm := make(map[byte]int)
	sm := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tm[t[i]]++
	}
	l, r, tl, tr := 0, 0, -1000001, 1000001
	have, need := 0, len(tm)
	for r < len(s) {
		sm[s[r]]++
		if sm[s[r]] == tm[s[r]] {
			have++
		}
		for have == need {
			// update the result here
			if (r - l + 1) <= (tr - tl) {
				tr = r
				tl = l
			}
			sm[s[l]]--
			if sm[s[l]] < tm[s[l]] {
				have--
			}
			l++
		}
		r++
	}
	if tl == -1000001 && tr == 1000001 {
		return ""
	}
	return s[tl : tr+1]
}

func isMatchedFreq(s1m, tm map[byte]int) bool {
	for k, v := range tm {
		if s1m[k] < v {
			return false
		}
	}
	return true
}

// bruet force solution
// func minWindow(s string, t string) string {
// 	if len(t) > len(s) {
// 		return ""
// 	}
// 	if strings.Contains(s, t) {
// 		// fmt.Println("from here")
// 		return t
// 	}
// 	tm := make(map[byte]int)
// 	sm := make(map[byte]int)
// 	for i := 0; i < len(t); i++ {
// 		tm[t[i]]++
// 	}
// 	l, r, tl, tr := 0, 0, -1000001, 1000001
// 	b := false
// 	for r < len(s) {
// 		sm[s[r]]++
// 		isMatched := isMatchedFreq(sm, tm)
// 		if isMatched {
// 			b = true
// 			for isMatched && l <= r {
// 				sm[s[l]]--
// 				l++
// 				isMatched = isMatchedFreq(sm, tm)
// 			}
// 			if (tr - tl) >= (r - (l - 1)) {
// 				tr = r
// 				tl = l - 1
// 			}
// 		}
// 		r++
// 	}
// 	res := ""
// 	if tl < 0 {
// 		tl = 0
// 	}
// 	if tr >= len(s) {
// 		tr = len(s) - 1
// 	}
// 	if tl == tr || !b {
// 		return ""
// 	}
// 	res = string(s[tl : tr+1])
// 	return res
// }
