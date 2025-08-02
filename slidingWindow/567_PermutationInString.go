package slidingwindow

// checkInclusion solved using variable size sliding window. Use map for storing frequency of
// characters in s1 compared to characters in current window. The time complexity is O(len(s2))
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freq := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		freq[s1[i]]++
	}

	count := 0
	left := 0

	for right := 0; right < len(s2); right++ {
		char := s2[right]

		// Decrease count in map
		if freq[char] > 0 {
			count++
		}
		freq[char]--

		// Maintain window size equal to len(s1)
		if right-left+1 > len(s1) {
			leftChar := s2[left]
			freq[leftChar]++
			if freq[leftChar] > 0 {
				count--
			}
			left++
		}

		// If all characters matched
		if count == len(s1) {
			return true
		}
	}

	return false
}
