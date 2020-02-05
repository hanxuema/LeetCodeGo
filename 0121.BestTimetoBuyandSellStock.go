package main

func maxProfit(prices []int) int {
    if prices == nil || len(prices) == 0{
        return 0
    }
    minPrice, maxProfit := prices[0], 0
    for i := 0; i < len(prices); i++{
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if (prices[i] - minPrice > maxProfit) {
            maxProfit = prices[i] - minPrice
        }
    }
    return maxProfit
}

// [7,1,5,3,6,4]
// [0,0,0,0]
// [8,7,6,5,3,2]
// [1,2,3]
// [1]