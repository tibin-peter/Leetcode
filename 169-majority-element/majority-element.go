func majorityElement(nums []int) int {
    hash:=make(map[int]int)

    for _,v:=range nums{
        hash[v]++
    }
    for _,v:=range nums{
        if hash[v]>len(nums)/2{
            return v
        }
    }
    return -1
}