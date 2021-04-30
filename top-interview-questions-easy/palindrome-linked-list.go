/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    var list []int

    node := head
    for node != nil {
        list = append(list, node.Val)
        node = node.Next
    }
   
    first := list[:len(list) / 2]
    second := list[len(list) /2:]
    
    for i:=0; i < len(first); i++ {
        if first[i] != second[len(second) - 1 - i] {
            return false
        }
    }
    
    return true
}
