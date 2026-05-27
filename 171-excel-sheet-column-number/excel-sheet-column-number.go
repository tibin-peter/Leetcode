func titleToNumber(columnTitle string) int {
    output:=0
    for _,v:= range columnTitle{
        numericValue:=int(v)-64 // for get the numeric value of a number
        output=output*26+numericValue
    }
    return output
}