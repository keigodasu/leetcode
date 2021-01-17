package main

func reverseList(head *ListNode) *ListNode {
    var current = head
    var prev *ListNode
    
    for current != nil {
        var next = current.Next
        current.Next = prev
        prev = current
        current = next
    }
    
    return prev
}


