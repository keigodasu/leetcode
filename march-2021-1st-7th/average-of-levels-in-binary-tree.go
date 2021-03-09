/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    var sum []int
    var count []int
    var results []float64
    sum2, count2 := average(root, 0, sum, count)
  
    for i := 0; i < len(sum2); i++ {
        results = append(results,(float64(sum2[i]) / float64(count2[i])))
    }
    
    return results
}

func average(node *TreeNode, i int, sum []int, count []int) ([]int, []int) {
    if node == nil {
        return sum, count
    }
    
    if i < len(sum)  {
        sum[i] = sum[i] + node.Val
        count[i]++
    } else {
        sum = append(sum, node.Val)
        count = append(count, 1)
    }
    
    a, b := average(node.Left, i + 1, sum, count)
    c, d := average(node.Right, i + 1, a, b)
    
    return c, d
}
