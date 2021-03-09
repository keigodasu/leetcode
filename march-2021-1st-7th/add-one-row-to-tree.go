/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, v int, d int) *TreeNode {
    insert(root, v, d, 1)
   
    return root
}

func insert(node *TreeNode, v int, d int, i int) {
    if node == nil {
        return 
    }
    
    if d == 1 {
        tempL := node.Left
        tempR := node.Right
        tempV := node.Val
        
        node.Val = v
        node.Left = &TreeNode {
            tempV,
            tempL,
            tempR, 
        } 
        node.Right = nil
        return 
    }    

    if d - 1 == i {
        //差し込み処理
        tempL := node.Left
        tempR := node.Right

        node.Left = &TreeNode {
            v,
            tempL,
            nil,
        }
        node.Right = &TreeNode {
            v,
            nil,
            tempR,
        }
        
        return
    }
    
    insert(node.Right, v, d, i+1)
    insert(node.Left, v, d, i+1)
    
}
