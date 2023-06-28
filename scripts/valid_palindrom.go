package scripts

import "strings"

func filterAlphaNumeric(s string) string {
	r := ""
	for _, v := range s {
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) || (v >= 48 && v <= 57) {
			r += strings.ToLower(string(v))
		}
	}
	return r
}

func IsPalindrome(s string) bool {
	s = filterAlphaNumeric(s)
	if len(s) == 0 {
		return true
	}
	l := len(s)
	for i := 0; i < l/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}
