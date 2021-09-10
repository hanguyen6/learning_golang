/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
* Use a queue to keep track of visiting node 
* Use variable to keep track the level of the tree 
* As queue is not empty
* Keep track the length of the queue as number of nodes in current level 
* Pop element(s) in current level from the queue
*   - Add their values to output list 
    - Check and insert left subtree and right subtree to queue 
* Go to the next level 
O(n) for time and space complexity 
*/
func levelOrder(root *TreeNode) [][]int {
    
    queue := make([]*TreeNode, 0)
    res := make([][]int, 0)
    
    if (root != nil) {
        queue = append(queue, root)
    }
    
    // Start from level 0 
    level := 0 
    
    // Adding element to queue 
    for len(queue) > 0 {
        
        // number of elements in the current level 
        level_len := len(queue)
        res = append(res, []int{})
        
        for i := 0; i < level_len; i++ {
            // Poping element from queue 
            elem := queue[0]
            queue = queue[1:]
            
            res[level] = append(res[level], elem.Val)
            
            if (elem.Left != nil) {
                queue = append(queue, elem.Left)
            }
        
            if (elem.Right != nil) {
                queue = append(queue, elem.Right)
            }       
        }
        // Go to the next level 
        level++    
        
    }
    
    return res
}
