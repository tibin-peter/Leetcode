/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function(nums1, m, nums2, n) {
    let newarr=[...nums1.slice(0,m),...nums2.slice(0,n)]
    newarr.sort((a,b)=>a-b)
    for(let i=0;i<nums1.length;i++){
       nums1[i]=newarr[i]
    }

    return nums1
};