/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
 /**
 * -- Recursive Solution
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
   
    if (root == nil) {
        return false
    }
    
    targetSum = targetSum - root.Val
    
    if (root.Left == nil && root.Right == nil) {
        return targetSum == 0
    } else {
        return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
    }
}
