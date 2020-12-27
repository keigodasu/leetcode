class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        recursiveHelper(root, result);
        
        return result;
    }
    
    public void recursiveHelper(TreeNode root, List<Integer> result)  {
        if(root != null) {
            if(root.left != null) {
                recursiveHelper(root.left, result);
            }
            result.add(root.val);
            if(root.right != null) {
                recursiveHelper(root.right, result);
            }
        }
    }
}
