/**
 * @param {number} n
 * @return {number[]}
 */
var sumZero = function(n) {
    const array=[]
    for(let i =1;i<=Math.floor(n/2);i++){
        array.push(i,-i)
    }
    if (n%2==1){
        array.push(0)
    }
    return array
    
};