func singleNumber(nums []int) []int {
    arr:=[]int{}
    hash:=make(map[int]int)

    for _,v:=range nums{
        hash[v]++
    }
    for i,v:= range hash{
        if v==1{
            arr=append(arr,i)
        }
    }
    return arr
}