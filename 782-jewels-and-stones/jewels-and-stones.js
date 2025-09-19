/**
 * @param {string} jewels
 * @param {string} stones
 * @return {number}
 */
var numJewelsInStones = function(jewels, stones) {
    let count=0
    for (index of jewels)
    for(let i=0;i<stones.length;i++){
    if(index==stones[i]){
        count++
    }
    }
    return count
};