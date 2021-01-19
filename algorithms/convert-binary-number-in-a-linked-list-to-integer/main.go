package convert_binary_number_in_a_linked_list_to_integer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	num := head.Val
	for head.Next != nil {
		num = (num << 1) | head.Next.Val
		head = head.Next
	}
	return num
}
