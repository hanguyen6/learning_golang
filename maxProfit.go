// Brute-force solution - O(n^2) time complexity 
// for each day, calculate profit/loss by selling stock in future 
// return max profit 
/*

func max(x int, y int) int {
    if (x >= y) {
        return x
    } else {
        return y
    }
}

func maxProfit(prices []int) int {
    m := 0 
    for i:=0; i < len(prices)-1; i++ {
        for j:=i+1; j < len(prices); j++ {
            m = max(m, prices[j] - prices[i])
        }
    }
    
    return m
    
}
*/

// Maintain a minPrice and a maxProfit
// Update minPrice, maxProfit as we iterate through the array 
// O(n) time complexity 
// O(1) space needed 

func maxProfit(prices []int) int {
    minPrice := prices[0]
    maxProfit := 0
    
    for i:=0; i < len(prices); i++ {
        if (prices[i] < minPrice) {
            minPrice = prices[i]
        }
        if (prices[i] - minPrice > maxProfit) {
            maxProfit = prices[i] - minPrice
        }
    }
    return maxProfit
}
