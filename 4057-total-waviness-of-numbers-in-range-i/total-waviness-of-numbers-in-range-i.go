func totalWaviness(num1 int, num2 int) int {
    count:=0
    for i:=num1;i<=num2;i++{
        numStr:=strconv.Itoa(i)
        for j:=1;j<=len(numStr)-2;j++{
            fmt.Println(numStr[j])
            if numStr[j]>numStr[j-1]&&numStr[j]>numStr[j+1]{
                count++
            }
            if numStr[j]<numStr[j-1]&&numStr[j]<numStr[j+1]{
                count++
            }
        }
    }
    return count
}