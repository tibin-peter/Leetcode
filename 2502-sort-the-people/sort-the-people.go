func sortPeople(names []string, heights []int) []string {
    n:=len(names)
    for i:=0;i<n;i++{
        for j:=0;j<n-i-1;j++{
            if heights[j]<heights[j+1]{
                heights[j],heights[j+1]=heights[j+1],heights[j]
                names[j],names[j+1]=names[j+1],names[j]
            }
        }
    }
    return names
}