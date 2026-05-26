func findDegrees(matrix [][]int) []int {
    n:=len(matrix)
    degree:=make([]int,n)
    for i:=0;i<n;i++{
        connections:=0
        for _,v:= range matrix[i]{
            connections+=v
        }
        degree[i]=connections
    }
    return degree
}
// func findDegrees(matrix [][]int) []int {
//     degree:=make([]int,len(matrix))
//     for i:=0;i<=len(matrix)-1;i++{
//         connections:=0
//         for j:=0;j<=len(matrix[i])-1;j++{
//             if matrix[i][j]==1{
//                 connections++
//             }
//         }
//         degree[i]=connections
//     }
//     return degree
// }

