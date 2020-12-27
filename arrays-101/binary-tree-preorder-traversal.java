class Solution {
    public List<Integer> preorderTraversal(TreeNode root) {
        LinkedList<TreeNode> stack = new LinkedList<>();
        LinkedList<Integer> results = new LinkedList<>();
        
        if(root == null) {
            return results;
        }
        
        stack.add(root);
        while(!stack.isEmpty()) {
            TreeNode node = stack.pollLast();
            results.add(node.val);
            if(node.right != null) {
                stack.add(node.right);
            } 
                
            if(node.left != null) {
                stack.add(node.left);
            }
        }
        
        return results;
    }
}
