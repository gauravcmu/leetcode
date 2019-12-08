func maxProfit(prices []int) int {
    profit_so_far := 0 
    max := 0 
    
    for i:=1; i< len(prices); i++ {
        if prices[i] > prices[i-1] {
            profit_so_far += prices[i] - prices[i-1]
        } 
    }
    
    if max == 0 {
        max = profit_so_far
    }
    
    return max 
}
