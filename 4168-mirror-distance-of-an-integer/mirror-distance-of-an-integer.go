func mirrorDistance(n int) int {
    reversed:=0
    originalValue:=n
    for n>0 {
        reversed= reversed*10+n%10
        n=n/10
    }
    output:=reversed-originalValue
    if output<0{
        output= -output
    }
    return output
}