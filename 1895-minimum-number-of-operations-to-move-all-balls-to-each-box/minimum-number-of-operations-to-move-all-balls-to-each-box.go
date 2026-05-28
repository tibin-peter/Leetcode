func minOperations(boxes string) []int {
    n:=len(boxes)
    operations:=make([]int,n)

    //for left to right
        leftBall:=0
        leftOpr:=0
        for i:=0;i<n;i++{
            operations[i]+=leftOpr
            if boxes[i]=='1'{
                leftBall++
            }
            leftOpr+=leftBall
        }
    // for Right-to-Left
        rightBalls := 0
        rightOps := 0
           for i := n - 1; i >= 0; i-- {
            operations[i] += rightOps
             if boxes[i] == '1' {
              rightBalls++
           }
            rightOps += rightBalls
       }
    return operations
}
// func minOperations(boxes string) []int {
//     n:=len(boxes)
//     operations:=make([]int,0,n)
//     for i:=0;i<n;i++{
//         count:=0
//         for index,v:=range boxes{
//             if v=='1'{
//                 distance:=index-i
//                 if distance<0{
//                     distance= -distance
//                 }
//                 count+=distance
//             }
//         }
//         operations=append(operations,count)
//     }
//     return operations
// }