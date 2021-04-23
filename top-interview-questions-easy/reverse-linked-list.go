/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    var list []int
    node := head
    
    for node != nil {
        list = append(list, node.Val)
        node = node.Next
    }
    
    newNpde := &ListNode{Val: list[0], Next: nil}
    list = list[1:]
    for _, v := range list {
        node := &ListNode{Val: v, Next: newNpde}
        newNpde = node
    } 
    
    return newNpde
}
