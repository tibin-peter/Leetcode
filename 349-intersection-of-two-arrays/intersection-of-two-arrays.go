func intersection(nums1 []int, nums2 []int) []int {
    hash1:=make(map[int]bool)
    for _,v:= range nums1{
        hash1[v]=true
    }
    hash2:=make(map[int]bool)
    for _,v:= range nums2{
        if hash1[v]{
            hash2[v]=true
        }
    }
    arr:=[]int{}
    for v:=range hash2{
        arr=append(arr,v)
    }
    return arr
}