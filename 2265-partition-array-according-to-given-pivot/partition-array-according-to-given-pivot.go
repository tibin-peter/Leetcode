func pivotArray(nums []int, pivot int) []int {
    leftArray:=[]int{}
    rightArray:=[]int{}
    equalArray:=[]int{}
    for _,v:=range nums{
        if v<pivot{
            leftArray=append(leftArray,v)
        }
        if v==pivot{
            equalArray=append(equalArray,v)
        }
        if v>pivot{
            rightArray=append(rightArray,v)
        }
    }
    outPut:=slices.Concat(leftArray,equalArray,rightArray)

    return outPut

}