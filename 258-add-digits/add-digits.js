/**
 * @param {number} num
 * @return {number}
 */
var addDigits = function(num) {
    let sum=num.toString().split("").map(Number).reduce((s,i)=>s+i,0)
    if(sum<10){
        return sum
    }
    return addDigits(sum)
    
};