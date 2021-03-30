/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
var fliped []int
var idx int
    var dfs func(root *TreeNode)
    
    if root == nil {
        return []int{-1}
    }
    

    
    dfs = func(root *TreeNode) {
        if root == nil {
            return
        }

        if root.Val != voyage[idx] {
            fliped = []int{-1}
            return
        }
        idx++

        if idx < len(voyage) && root.Left != nil && root.Left.Val != voyage[idx] {
            fliped = append(fliped, root.Val)
            dfs(root.Right)
            dfs(root.Left)
        } else{
            dfs(root.Left)
            dfs(root.Right)
        }
    }
    
    dfs(root)
    if fliped != nil && fliped[0] == -1 {
        return []int{-1}
    }
    
    return fliped
}
