/**
 * @param {number[]} nums
 * @param {number} original
 * @return {number}
 */
var findFinalValue = function(nums, original) {
    let result=original
    for(let i=0;i<nums.length;i++){
        if(nums[i]==result){
            result=result*2
        return findFinalValue(nums,result)
        }
    }
    return result
};