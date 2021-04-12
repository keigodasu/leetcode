type Queue []*TreeNode

func (q *Queue) Push(val *TreeNode) {
    *q = append(*q, val)
}

func (q *Queue) Pop() *TreeNode {
    popped := (*q)[0]
    (*q) = (*q)[1:len(*q)]
    return popped
}

func (q *Queue) isEmpty() bool {
    return len(*q) == 0
}

func deepestLeavesSum(root *TreeNode) int {
    sum := 0
    if root == nil {
        return sum
    }
    
    q := &Queue{root}
    
    for !q.isEmpty() {
        sum = 0
        l := len(*q) 
        for i := 0; i < l; i++ {
            pop := q.Pop()
            sum += pop.Val
            
            if pop.Left != nil {
                q.Push(pop.Left)
            }
            if pop.Right != nil {
                q.Push(pop.Right)
            }
        }
    }
    
    return sum
}
