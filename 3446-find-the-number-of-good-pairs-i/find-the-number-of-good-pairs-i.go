func numberOfPairs(nums1 []int, nums2 []int, k int) int {
    count:=0
    for i:=0;i<len(nums2);i++{
        for j:=0;j<len(nums1);j++{
            if nums1[j]%(nums2[i]*k)==0{
                count++
            }
        }
    }
    return count
}