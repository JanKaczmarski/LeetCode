package twoPointers

func twoSum(numbers []int, target int) []int {
	lp, rp := 0, len(numbers)-1

	for lp < rp {
		s := numbers[lp] + numbers[rp]
		if s == target {
			return []int{lp + 1, rp + 1}
		} else if s < target {
			lp++
		} else {
			rp--
		}
	}
	panic("No result was found!")
}
