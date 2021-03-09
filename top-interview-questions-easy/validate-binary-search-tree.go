/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return isValidate(root, nil, nil)
}

func isValidate(node *TreeNode, low *int, high *int) bool {
    if node == nil {
        return true
    }
    
    v := &node.Val
    
    if (low != nil && node.Val <= *low) || (high != nil && node.Val >= *high) {
        return false
    }
    
    return isValidate(node.Right, v, high) && isValidate(node.Left, low, v)
}
