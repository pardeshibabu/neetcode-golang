package scripts

import (
	"fmt"
	"strings"
)

func isMatchedFreq(s1m, tm map[byte]int) bool {
	// fmt.Println("s1m", s1m)
	// fmt.Println("tm", tm)
	// fmt.Println("##################3")
	for k, v := range tm {
		if s1m[k] < v {
			// fmt.Println("from here")
			return false
		}
	}
	return true
}

func MinWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	if strings.Contains(s, t) {
		return t
	}
	tm := make(map[byte]int)
	sm := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tm[t[i]]++
	}
	requiredCount := len(t)
	mc := 0
	l, r, tl, tr := 0, 0, -1000001, 1000001
	// b := false
	for r < len(s) {
		sm[s[r]]++
		if sm[s[r]] >= tm[s[r]] {
			mc++
		}
		fmt.Println(sm)
		fmt.Println(tm)
		fmt.Println(mc, requiredCount, l, r)
		for mc == requiredCount && l <= r {
			sm[s[l]]--
			if sm[s[l]] < tm[s[l]] {
				mc--
			}
			l++
		}
		ind := l - 1
		if l <= 0 {
			ind = 0
		}
		if sm[s[ind]]+1 == tm[s[ind]] && mc+1 == requiredCount && (tr-tl) >= (r-(l-1)) {
			tl = ind
			tr = r
		}
		fmt.Println(l, r, tl, tr)
		fmt.Println("#############")
		r++
		// 	isMatched := isMatchedFreq(sm, tm)
		// 	if isMatched {
		// 		b = true
		// 		for isMatched && l <= r {
		// 			sm[s[l]]--
		// 			l++
		// 			isMatched = isMatchedFreq(sm, tm)
		// 		}
		// 		// fmt.Println(l, r)
		// 		// fmt.Println("::1111::",tl, tr)
		// 		if (tr - tl) >= (r - (l - 1)) {
		// 			tr = r
		// 			tl = l - 1
		// 		}
		// 		// fmt.Println("::222tl:::",tl, tr)
		// 	}
		// r++
		// }
		// res := ""
		// if tl < 0 {
		// 	tl = 0
		// }
		// if tr >= len(s) {
		// 	tr = len(s) - 1
		// }
		// if tl == tr || !b {
		// 	return ""
	}
	res := ""
	res = string(s[tl : tr+1])
	return res
	return ""
}
