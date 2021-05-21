/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l := 1
    r := 1
    
    left := recur(root.Left, l)
    right := recur(root.Right, r)
    
    return int(math.Max(float64(left), float64(right)))
}

func recur(node *TreeNode, num int) int {
    if node == nil {
        return num
    } else {
        num++
    } 
    
    left := recur(node.Left, num)
    right := recur(node.Right, num)
    
    return int(math.Max(float64(left), float64(right)))
}
