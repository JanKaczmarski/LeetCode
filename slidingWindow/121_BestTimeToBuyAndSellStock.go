package slidingwindow

func maxProfit(prices []int) int {
	// two-pointers: first points to left (day to buy) 2nd pointer goes through the array
	// when smaller value then first pointer points to is met we swap to it. In each iteration we check
	// if we have better result with current left and right pointer combination
	l, res := 0, 0

	for r := 1; r < len(prices); r++ {
		if prices[r]-prices[l] > res {
			res = prices[r] - prices[l]
		}
		if prices[r] < prices[l] {
			l = r
		}
	}

	return res
}
