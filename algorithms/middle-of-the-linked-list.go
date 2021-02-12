/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    total := 1
    
    ptr := head
    for ptr.Next != nil {
        total++
        ptr = ptr.Next
    }
    fmt.Println(total)
    
    ptr = head
    for i:=0; i < total / 2 ; i++ {
        ptr = ptr.Next
    }
    
    return ptr
}
