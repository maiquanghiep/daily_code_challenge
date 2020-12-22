package sellstock

import (
	"math"
	"strconv"
)

/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.
*/

/*
	Loop from the end of the slice, find the Max number each time and calculate the max profit

*/

func findMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
func maxProfit1(prices []int) int {
	var max, profix int
	for i := len(prices) - 2; i >= 0; i-- {
		max = findMax(max, prices[i+1])
		if max-prices[i] > profix {
			profix = max - prices[i]
		}
	}
	return profix
}

/*
Say you have an array prices for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
*/
func maxProfit2(prices []int) int {
	min := math.MaxInt32
	var profit, result int
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			profit = 0
		} else {
			if prices[i]-min > profit {
				result -= profit
				profit = prices[i] - min
				result += profit
			} else {
				min = prices[i]
				profit = 0
			}
		}
	}
	return result
}

/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most two transactions.

Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
*/
func findProfit3(prices []int, pos int, t int, bought int, memo map[string]int) int {
	if pos == len(prices) || t == 0 {
		return 0
	}

	action := strconv.Itoa(bought) + "," + strconv.Itoa(t) + "," + strconv.Itoa(pos)
	if v, ok := memo[action]; ok {
		return v
	}
	result := findProfit3(prices, pos+1, t, bought, memo)
	if bought == 1 {
		result = findMax(result, findProfit3(prices, pos+1, t-1, 0, memo)+prices[pos])
	} else {
		result = findMax(result, findProfit3(prices, pos+1, t, 1, memo)-prices[pos])
	}
	memo[action] = result
	return result
}
func maxProfit3(prices []int) int {
	memo := make(map[string]int)
	result := findProfit3(prices, 0, 2, 0, memo)
	return result
}
