/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
   firstPtr := head
    
    jump := k
    for jump > 1 {
        firstPtr = firstPtr.Next
        jump--
    }
    
    jumpPtr := firstPtr
    secondPtr := head
    
    for jumpPtr.Next != nil {
        jumpPtr = jumpPtr.Next
        secondPtr = secondPtr.Next
    }
    
    firstPtr.Val, secondPtr.Val = secondPtr.Val, firstPtr.Val
    
    return head
}
