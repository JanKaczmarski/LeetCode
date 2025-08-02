package slidingwindow

func characterReplacement(s string, k int) int {
	track := make([]int, 26)
	l, maxFreq, result := 0, 0, 0

	for r := 0; r < len(s); r++ {
		chIdx := int(s[r] - 'A')

		// Update how many char s[r] was seen
		track[chIdx]++

		// Update maxFreq
		if track[chIdx] > maxFreq {
			maxFreq = track[chIdx]
		}

		// adjust window size
		if (r-l+1)-maxFreq > k {
			track[int(s[l]-'A')]--
			l++
		}

		if r-l+1 > result {
			result = r - l + 1
		}
	}

	return result
}
