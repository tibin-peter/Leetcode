/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findDuplicates = function(nums) {
    const empty=new Set()
  return nums.filter(n=>empty.size===empty.add(n).size)
    
};