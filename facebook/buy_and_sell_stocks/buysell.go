func maxProfit(prices []int) int {
    profit_so_far := 0 
    max := 0 
    
    for i:=1; i< len(prices); i++ {
        profit_so_far += prices[i] - prices[i-1]
        
        //if profit so far goes negative, then reset it. (not just on it going below previous profit_so_far)
        if profit_so_far < 0 {
            profit_so_far = 0 
        }
        
        if profit_so_far > max {
            max = profit_so_far
        }
    }
    return max 
}
