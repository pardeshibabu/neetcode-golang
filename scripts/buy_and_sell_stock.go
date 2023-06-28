package scripts

import "fmt"

func MaxProfit(prices []int) int {
	l, r, maxProfit := 0, 0, 0
	for r < len(prices) {
		fmt.Println("r:::", r)
		tempmax := (prices[r] - prices[l])
		fmt.Println(tempmax, maxProfit)
		if maxProfit < tempmax {
			maxProfit = tempmax
		}
		if prices[l] > prices[r] {
			l = r
		}
		r++
	}
	return maxProfit
}
