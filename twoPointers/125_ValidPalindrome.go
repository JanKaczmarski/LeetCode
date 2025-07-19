package twoPointers

import (
	"strings"
)

func isPalindrome(s string) bool {
	lp, rp := 0, len(s)-1
	s = strings.ToLower(s)

	for ; lp < rp; lp, rp = lp+1, rp-1 {
		// Skip non-alphanumeric values
		for lp < rp && !isAlphaNumeric(s[lp]) {
			lp++
		}
		for lp < rp && !isAlphaNumeric(s[rp]) {
			rp--
		}

		// No more non-alphanumeric chars to compare
		if lp > rp {
			return true
		}

		if s[lp] != s[rp] {
			return false
		}
	}
	return true
}

func isAlphaNumeric(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}
