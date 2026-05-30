func pivotArray(nums []int, pivot int) []int {
    n:=len(nums)
    result:=make([]int,n)
    left:=0
    right:=n-1
    for i:=0;i<n;i++{
        if nums[i]<pivot{
            result[left]=nums[i]
            left++
        }
        j:=n-1-i
        if nums[j]>pivot{
            result[right]=nums[j]
            right--
        }
    }
    for left <= right { //for pushing the qual ones
        result[left] = pivot
        left++
    }
    return result
}
// func pivotArray(nums []int, pivot int) []int {
//     leftArray:=[]int{}
//     rightArray:=[]int{}
//     equalArray:=[]int{}
//     for _,v:=range nums{
//         if v<pivot{
//             leftArray=append(leftArray,v)
//         }
//         if v==pivot{
//             equalArray=append(equalArray,v)
//         }
//         if v>pivot{
//             rightArray=append(rightArray,v)
//         }
//     }
//     outPut:=slices.Concat(leftArray,equalArray,rightArray)

//     return outPut

// }