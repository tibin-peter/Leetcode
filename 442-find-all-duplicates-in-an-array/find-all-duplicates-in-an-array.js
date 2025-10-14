/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findDuplicates = function(nums) {
//     const empty=new Set()
//   return nums.filter(n=>empty.size===empty.add(n).size)
const empty=new Set()
  const duplicate =[]
  for (const n of nums){
      empty.has(n)?duplicate.push(n):empty.add(n)
  }
  return duplicate
    
};