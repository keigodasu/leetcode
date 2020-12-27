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

class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> result = new ArrayList<>();
        Stack<TreeNode> stack = new Stack<>();
        TreeNode curr = root;
        while(curr != null || !stack.isEmpty()) {
            while(curr != null) {
                stack.push(curr);
                curr = curr.left;
            }
            curr = stack.pop();
            result.add(curr.val);
            curr = curr.right;
        }
        return result;
    }
}
