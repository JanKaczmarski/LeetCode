package slidingwindow

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]bool)
	// left border of dynamic-len sliding window
	l := 0
	longest := 0

	for r := 0; r < len(s); r++ {
		// shrink window to remove duplicates
		for set[s[r]] {
			set[s[l]] = false
			l++
		}

		set[s[r]] = true

		if r-l+1 > longest {
			longest = r - l + 1
		}
	}

	return longest
}
