package scripts

import "fmt"

func CheckInclusion(s1 string, s2 string) bool {
	s1m := make(map[int]int)
	s2m := make(map[int]int)
	for i := 0; i < len(s1); i++ {
		s1m[int(s1[i]-'a')]++
		s2m[int(s2[i]-'a')]++
	}
	matches := 0
	for i := 0; i < 26; i++ {
		if s1m[i] == s2m[i] {
			matches++
		}
	}
	l := 0
	fmt.Println(s1m)
	fmt.Println(s2m)
	for r := len(s1); r < len(s2); r++ {
		if matches == 26 {
			return true
		}
		ind := int(s2[r] - 'a')
		s2m[ind]++
		if s2m[ind] == s1m[ind] {
			matches++
		} else if s2m[ind] == s1m[ind]+1 {
			matches--
		}
		ind1 := int(s2[l] - 'a')
		s2m[ind1]--
		if s2m[ind1] == s1m[ind1] {
			matches++
		} else if s2m[ind1] == s1m[ind1]-1 {
			matches--
		}
		l++
	}
	return false
}
