func gcdOfOddEvenSums(n int) int {
    sumOdd:=0
    sumEven:=0
    for i:=1;i<=n*2;i++{
        if i%2==0{
            sumEven+=i
        }else{
            sumOdd+=i
        }
    }
    return sumEven-sumOdd
    
}