func intersect(nums1 []int, nums2 []int) []int {
    arr:=[]int{}
    if len(nums1)>len(nums2){
        nums1,nums2=nums2,nums1
    }
    hash:=make(map[int]int)
    for _,v:= range nums1{
        hash[v]++
    }
    for _,v:=range nums2{
        if hash[v]>0{
            arr=append(arr,v)
            hash[v]--
        }
    }
    return arr
}