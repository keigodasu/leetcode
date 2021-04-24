/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    var list []int
    node := l1
    
    for node != nil {
        list = append(list, node.Val)
        node = node.Next
    }
    
    node = l2
    for node != nil {
        list = append(list, node.Val)
        node = node.Next
    }
    
    sort.Ints(list)
    fmt.Println(list)
    if len(list) == 0 {
        return nil
    }
    if len(list) == 1 {
        return &ListNode{Val: list[0], Next:nil}
    }
    
    nnode := &ListNode{Val: list[len(list) - 1], Next: nil}
    for i := len(list) - 2; i >= 0; i-- {
        node := &ListNode{Val: list[i], Next: nnode}
        nnode = node
    }
    
    return nnode
}
