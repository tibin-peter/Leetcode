/**
 * @param {number[]} nums
 * @return {number}
 */
var findNumbers = function(nums) {
    let result=nums.map((i)=>i.toString().split("").length)
    let count=0
    result.forEach((i)=>{
        if(i%2==0){
            count++
        }
    })
    return count
};