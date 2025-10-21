/**
 * @param {number[]} nums
 * @return {number}
 */
var missingNumber = function(nums) {
    let cursum=nums.reduce((sum,i)=>sum+i,0)
    let expsum=0
    for(let i=1;i<=nums.length;i++){
        expsum+=i
    }
    return expsum-cursum
};