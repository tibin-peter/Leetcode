/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let maximumprofit=0
    let  smallprice=Infinity
    for(let i = 0;i<prices.length;i++){
        if(prices[i]<smallprice){
            smallprice=prices[i]
        }
        let profit=prices[i]-smallprice
       if(profit>maximumprofit){
           maximumprofit=profit
    }
    }
    return  maximumprofit
    
};