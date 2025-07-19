package twoPointers

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	// SORT the array!!!
	res := make([][]int, 0)
	n := len(nums)

	fmt.Printf("Nums: %v\n", nums)

	// shrink window by reducing k
	for i := 0; i+2 < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, n-1
		for j < k {
			s := nums[i] + nums[j] + nums[k]

			if s == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				// jump to avoid putting single triplet multiple times
				j++
				for nums[j] == nums[j-1] && j < k {
					j++
				}
			} else if s > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return res
}
