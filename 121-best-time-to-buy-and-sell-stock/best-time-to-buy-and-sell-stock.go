func maxProfit(prices []int) int {
    minPrice:=prices[0]
    profit :=0
    for i:=range prices{
        if len(prices)==0{
            return 0
        }
        if minPrice>prices[i]{
            minPrice=prices[i]
        }else if prices[i]-minPrice>profit{
            profit = prices[i]-minPrice
        }
    }
    return profit
}