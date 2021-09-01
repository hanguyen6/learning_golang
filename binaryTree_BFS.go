// https://leetcode.com/problems/binary-tree-level-order-traversal/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// O(n) for time & space complexity 
func levelOrder(root *TreeNode) [][]int {
    var levels = [][]int{}
    if (root == nil) {
        return levels
    }
    
    queue := []TreeNode{}
    queue = append(queue, *root)
    level := 0 
    
    for len(queue) != 0 {
        
        levels = append(levels, []int{})
        level_len := len(queue)
        // check elements in current level 
        for i:=0; i < level_len; i++ {
            node := queue[0]
            queue = queue[1:]
            levels[level] = append(levels[level], node.Val)
            
            // Add child nodes of current levels to queue 
            if (node.Left != nil) {
                queue = append(queue, *node.Left)
            }
            if (node.Right != nil) {
                queue = append(queue, *node.Right)
            }
        }
        level++

    }
    
    return levels
    
   
}
