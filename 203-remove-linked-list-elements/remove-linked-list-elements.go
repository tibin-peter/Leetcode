/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    sample:=&ListNode{Next:head}
    current:=sample
    for current.Next!=nil{
        if current.Next.Val==val{
        current.Next=current.Next.Next
    }else{
        current=current.Next
    }
    }
    return sample.Next
}