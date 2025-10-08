/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    let unique=new Set(nums).size
    if(nums.length==unique){
        return false
    }else{
        return true
    }
};