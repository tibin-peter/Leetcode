func minimumAbsDifference(arr []int) [][]int {
    pairs:=[][]int{}
    n:=len(arr)
    sort.Slice(arr,func(a,b int)bool{
        return arr[a]<arr[b]
    })
    minimumDifference:=arr[1]-arr[0]
    for i:=1;i<n;i++{
        dif:=arr[i]-arr[i-1]
        if dif<minimumDifference{
            minimumDifference=dif
        }
    }
    for i:=0;i<n;i++{
        j:=i+1
        if j<n&&arr[j]-arr[i]==minimumDifference{
            pair:=[]int{arr[i],arr[j]}
            pairs=append(pairs,pair)
        }
    }
    return pairs
}