class Solution {
    int count = 0;
    
    public boolean isValidPart(TreeNode node, int val) {
        if(node == null) return true;
        if(!isValidPart(node.left, node.val) | !isValidPart(node.right, node.val)) return false;
        count++;
        return node.val == val;
    }
    
    public int countUnivalSubtrees(TreeNode root) {
        isValidPart(root, 0);
        return count;
    }
}
